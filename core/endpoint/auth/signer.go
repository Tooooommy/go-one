package auth

import (
	"context"
	"errors"
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
)

var (
	ErrInvalidClaims = errors.New("claims is invalid")
)

func KitSigner(secret []byte, claims jwt.MapClaims) endpoint.Middleware {
	kid := uuid.Must(uuid.NewUUID()).String()
	return kitjwt.NewSigner(kid, secret, jwt.SigningMethodHS256, claims)
}

func Signer(secret []byte, timeout time.Duration) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = next(ctx, request) // 先执行
			if err != nil {
				return nil, err
			}
			claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(jwt.MapClaims)
			if !ok {
				return response, nil
			}

			// 生成
			buildClaims(claims, JWTIssuer, "go-one")
			buildClaims(claims, JWTSubject, "go-one")
			buildClaims(claims, JWTAudience, "go-one")
			buildClaims(claims, JWTIssueAt, time.Now().Unix())
			buildClaims(claims, JWTExpire, time.Now().Add(timeout).Unix())
			buildClaims(claims, JWTNotBefore, time.Now().Unix())
			buildClaims(claims, JWTId, uuid.Must(uuid.NewUUID()).String())
			signer := KitSigner(secret, claims)

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

func buildClaims(claims jwt.MapClaims, key string, val interface{}) {
	if _, ok := claims[key]; !ok {
		claims[key] = val
	}
}
