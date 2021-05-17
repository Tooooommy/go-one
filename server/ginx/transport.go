package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type (
	DecodeFunc  func(*gin.Context) (interface{}, error)
	EncodeFunc  func(*gin.Context, interface{}) error
	ErrorFunc   func(*gin.Context, error)
	DecodeFuncs []DecodeFunc
	Options     func(*ErrHandler)

	ErrHandler struct {
		encodeErrHandler   ErrorFunc
		decodeErrHandler   ErrorFunc
		endpointErrHandler ErrorFunc
	}
)

func defaultHandleErr(ctx *gin.Context, err error) {
	_ = ctx.AbortWithError(http.StatusBadRequest, err)
	return
}

func NewService(e endpoint.Endpoint, decode DecodeFuncs, encode EncodeFunc, options ...Options) gin.HandlerFunc {
	errHandler := &ErrHandler{
		encodeErrHandler:   defaultHandleErr,
		decodeErrHandler:   defaultHandleErr,
		endpointErrHandler: defaultHandleErr,
	}
	for _, option := range options {
		option(errHandler)
	}

	return func(ctx *gin.Context) {
		// 解码
		var request interface{}
		var err error
		for _, dec := range decode {
			request, err = dec(ctx)
			if err != nil {
				errHandler.decodeErrHandler(ctx, err)
				return
			}
		}

		// 逻辑
		response, err := e(ctx, request)
		if err != nil {
			errHandler.endpointErrHandler(ctx, err)
			return
		}

		// 解码
		err = encode(ctx, response)
		if err != nil {
			errHandler.encodeErrHandler(ctx, err)
			return
		}
	}
}
