package recovery

import (
	"context"
	"github.com/Tooooommy/go-one/core/response/jsonx"
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
				case *jsonx.Response:
					response = res
				case error:
					response = jsonx.RawJSON(jsonx.Failure, res.Error(), nil)
				case string:
					response = jsonx.RawJSON(jsonx.Failure, res, nil)
				case int:
					response = jsonx.RawJSON(res, "no response", nil)
				default:
					zapx.Error(ctx).Any("Recovery Panic", res).Msg(string(debug.Stack()))
				}
			}()
			response, err = next(ctx, request)
			return
		}
	}
}
