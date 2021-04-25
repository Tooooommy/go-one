package x

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

/*
from go-zero
*/

const (
	jwtAudience  = "aud"
	jwtExpire    = "exp"
	jwtId        = "jti"
	jwtIssueAt   = "iat"
	jwtIssuer    = "iss"
	jwtNotBefore = "nbf"
	jwtSubject   = "sub"
	jwtRealIp    = "rip"
	jwtDevice    = "dvc"
	JwtAuthToken = "JWT_AUTH_TOKEN" // TOKEN
)

var (
	xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP       = http.CanonicalHeaderKey("X-Real-IP")
)

var (
	ErrInvalidToken  = errors.New("token is invalid")
	ErrInvalidClaims = errors.New("claims is invalid")
	ErrInvalidRealIp = errors.New("real ip is invalid")
	ErrInvalidDevice = errors.New("device is invalid")
)

type (
	TokenConfig struct {
		Secret    string `json:"secret"`
		PreSecret string `json:"pre_secret"`
	}

	ParseOption func(parser *TokenParser)

	// 解析器
	TokenParser struct {
		cfg          TokenConfig
		resetTime    time.Time
		restDuration time.Duration
		history      sync.Map
		validRealIp  bool // 校验IP
		validDevice  bool // 校验Device
	}

	BuildOption func(builder *TokenBuilder)
	// 生成器
	TokenBuilder struct {
		secret  string
		timeout time.Duration
	}
)

func NewTokenParser(cfg TokenConfig, opts ...ParseOption) *TokenParser {
	parser := &TokenParser{
		cfg:          cfg,
		resetTime:    time.Now(),
		restDuration: 24 * time.Hour,
		history:      sync.Map{},
		validRealIp:  false,
		validDevice:  false,
	}
	for _, opt := range opts {
		opt(parser)
	}
	return parser
}

func WithRestDuration(duration time.Duration) ParseOption {
	return func(parser *TokenParser) {
		parser.restDuration = duration
	}
}

func (p *TokenParser) ParseToken(r *http.Request) (*jwt.Token, error) {
	var token *jwt.Token
	var err error
	if len(p.cfg.PreSecret) > 0 {
		count := p.loadCount(p.cfg.Secret)
		preCount := p.loadCount(p.cfg.PreSecret)

		first := p.cfg.Secret
		second := p.cfg.PreSecret
		if count < preCount {
			first = p.cfg.PreSecret
			second = p.cfg.Secret
		}
		token, err = p.doParseToken(r)
		if err != nil {
			token, err = p.doParseToken(r)
			if err != nil {
				return nil, err
			}
			p.incrementCount(second)
		} else {
			p.incrementCount(first)
		}
	} else {
		token, err = p.doParseToken(r)
		if err != nil {
			return nil, err
		}
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, JwtAuthToken, token.Raw)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	if err = claims.Valid(); err != nil {
		return nil, err
	}

	if err = p.VerifyRealIp(r.RemoteAddr, claims); err != nil {
		return nil, err
	}

	if err = p.VerifyDevice(r.UserAgent(), claims); err != nil {
		return nil, err
	}

	for k, v := range claims {
		switch k {
		case jwtAudience, jwtExpire, jwtRealIp, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
		default:
			// 时间戳
			ctx = context.WithValue(ctx, k, v)
		}
	}
	r.WithContext(ctx)
	return token, nil
}

func (p *TokenParser) VerifyDevice(device string, claims jwt.MapClaims) error {
	if p.validDevice {
		if vDevice, ok := claims[jwtDevice]; ok && device == vDevice {
			return nil
		} else {
			return ErrInvalidDevice
		}
	}
	return nil
}

func (p *TokenParser) VerifyRealIp(rip string, claims jwt.MapClaims) error {
	if p.validRealIp {
		if vip, ok := claims[jwtRealIp]; ok && rip == vip {
			return nil
		} else {
			return ErrInvalidRealIp
		}
	}
	return nil
}

func (p *TokenParser) loadCount(secret string) uint64 {
	value, ok := p.history.Load(secret)
	if ok {
		return *value.(*uint64)
	}
	return 0
}

func (p *TokenParser) incrementCount(secret string) {
	if p.resetTime.Add(p.restDuration).Before(time.Now()) {
		p.history.Range(func(key, value interface{}) bool {
			p.history.Delete(key)
			return true
		})
	}
	value, ok := p.history.Load(secret)
	if ok {
		atomic.AddUint64(value.(*uint64), 1)
	} else {
		var count uint64 = 1
		p.history.Store(secret, &count)
	}
}

func (p *TokenParser) doParseToken(r *http.Request) (*jwt.Token, error) {
	return request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(p.cfg.Secret), nil
		}, request.WithParser(&jwt.Parser{
			UseJSONNumber: true,
		}))
}

func NewTokenBuilder(opts ...BuildOption) *TokenBuilder {
	builder := &TokenBuilder{
		timeout: 24 * time.Hour,
	}
	for _, opt := range opts {
		opt(builder)
	}
	return builder
}

func (b *TokenBuilder) WithSecret(secret string) BuildOption {
	return func(builder *TokenBuilder) {
		builder.secret = secret
	}
}

func (b *TokenBuilder) WithTimeout(timeout time.Duration) BuildOption {
	return func(builder *TokenBuilder) {
		builder.timeout = timeout
	}
}

func (b *TokenBuilder) BuildToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(b.secret)
	return tokenStr, err
}
