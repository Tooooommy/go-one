package auth

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

func KitParser(secret []byte) endpoint.Middleware {
	return kitjwt.NewParser(func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}, jwt.SigningMethodHS256, kitjwt.MapClaimsFactory)
}

func Parser(secret []byte, preSecret []byte) endpoint.Middleware {
	parser := KitParser(secret)
	preParser := KitParser(preSecret)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = parser(next)(ctx, request)
			if err != nil {
				response, err = preParser(next)(ctx, request)
			}
			return
		}
	}
}
