package recovery

import (
	"context"
	"github.com/Tooooommy/go-one/core/httpx"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/go-kit/kit/endpoint"
	"runtime/debug"
)

func Recovery() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				result := recover()
				switch res := result.(type) {
				case *httpx.JSONResponse:
					response = res
				case error:
					response = httpx.RawJSON(httpx.Failure, res.Error(), nil)
				case string:
					response = httpx.RawJSON(httpx.Failure, res, nil)
				case int:
					response = httpx.RawJSON(res, "no response", nil)
				default:
					zapx.Error().Any("Recovery Panic", res).Msg(string(debug.Stack()))
				}
			}()
			response, err = next(ctx, request)
			return
		}
	}
}
