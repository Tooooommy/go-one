package auth

import (
	"context"
	"errors"
	"github.com/Tooooommy/go-one/server/config"
	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"time"
)

const (
	JWTAudience  = "aud"
	JWTExpire    = "exp"
	JWTId        = "jti"
	JWTIssueAt   = "iat"
	JWTIssuer    = "iss"
	JWTNotBefore = "nbf"
	JWTSubject   = "sub"
	JWTRealIp    = "rip"
	JWTDevice    = "dvc"
)

var (
	ErrInvalidToken     = errors.New("token is invalid")
	ErrInvalidClaims    = errors.New("claims is invalid")
	ErrInvalidRealIp    = errors.New("real ip is invalid")
	ErrInvalidDevice    = errors.New("device is invalid")
	ErrInvalidKidHeader = errors.New("kid header is invalid")
	ErrInvalidRefresh   = errors.New("refresh token is disable")
)

const (
	JWTClientRealIp = "JWT_CLIENT_REAL_IP"
	JWTClientDevice = "JWT_CLIENT_DEVICE"
	JWTKidHeader    = "JWT_KID_HEADER"
)

// NewJwtParser: 解析JWT数据
func NewJwtParser(cfg config.JwtConfig) endpoint.Middleware {
	return parse([]byte(cfg.Secret), cfg.ValidRealIp, cfg.ValidDevice)
}

// NewJwtSigner: 生成JWTTOken
func NewJwtSigner(cfg config.JwtConfig) endpoint.Middleware {
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 24
	}
	return signe([]byte(cfg.Secret), timeout)
}

func buildClaims(claims jwt.MapClaims, key string, val interface{}) {
	if _, ok := claims[key]; !ok {
		claims[key] = val
	}
}

// NewJwtReSigner: 重新签发JWTToken, 修改secret
func NewJwtReSigner(cfg config.JwtConfig) endpoint.Middleware {
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 24
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		parser := parse([]byte(cfg.PreSecret), cfg.ValidRealIp, cfg.ValidDevice)
		return parser(func(ctx context.Context, request interface{}) (response interface{}, err error) {
			return signe([]byte(cfg.Secret), timeout)(next)(ctx, request)
		})
	}
}

func parse(secret []byte, validRealIp, validDevice bool) endpoint.Middleware {
	parser := kitjwt.NewParser(func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}, jwt.SigningMethodHS256, kitjwt.MapClaimsFactory)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return parser(func(ctx context.Context, request interface{}) (response interface{}, err error) {
			claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(jwt.MapClaims)
			if !ok {
				return nil, ErrInvalidClaims
			}
			for k, v := range claims {
				switch k {
				case JWTAudience, JWTExpire, JWTId, JWTIssueAt, JWTIssuer, JWTNotBefore, JWTSubject:

				case JWTRealIp:
					if validRealIp && ctx.Value(JWTClientRealIp) != v {
						return nil, ErrInvalidRealIp
					}

				case JWTDevice:
					if validDevice && ctx.Value(JWTClientDevice) != v {
						return nil, ErrInvalidDevice
					}

				default:
					// 时间戳
					ctx = context.WithValue(ctx, k, v)
				}
			}
			return next(ctx, request)
		})
	}
}

func signe(secret []byte, timeout int) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = next(ctx, request) // 先执行
			if err != nil {
				return nil, err
			}
			// 获取->
			kid, ok := ctx.Value(JWTKidHeader).(string)
			if !ok {
				return nil, ErrInvalidKidHeader
			}
			claims, ok := ctx.Value(response).(jwt.MapClaims)
			if !ok {
				return nil, ErrInvalidClaims
			}

			// 生成
			buildClaims(claims, JWTIssuer, "go-one")
			buildClaims(claims, JWTSubject, "go-one")
			buildClaims(claims, JWTAudience, "go-one")
			buildClaims(claims, JWTIssueAt, time.Now().Unix())
			buildClaims(claims, JWTExpire, time.Now().Add(time.Duration(timeout)*time.Hour).Unix())
			buildClaims(claims, JWTDevice, ctx.Value(JWTClientDevice))
			buildClaims(claims, JWTRealIp, ctx.Value(JWTClientRealIp))
			buildClaims(claims, JWTNotBefore, time.Now().Unix())
			buildClaims(claims, JWTId, uuid.Must(uuid.NewUUID()).String())
			signer := kitjwt.NewSigner(kid, secret, jwt.SigningMethodHS256, claims)

			// 转换成Token
			return signer(func(ctx context.Context, request interface{}) (interface{}, error) {
				tokenString, ok := ctx.Value(kitjwt.JWTTokenContextKey).(string)
				if !ok {
					return nil, ErrInvalidClaims
				}
				return tokenString, nil
			})(ctx, request)
		}
	}
}
