package ginx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"mime/multipart"
)

type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type JSONEndpoint func(ctx context.Context, request interface{}) (*JSONResponse, error)

func JSONToEndpoint(e JSONEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return e(ctx, request)
	}
}

var ErrFileHeaderInvalid = errors.New("file header is invalid")

type FileEndpoint func(ctx context.Context, fs ...*multipart.FileHeader) (*JSONResponse, error)

func FileToEndpoint(e FileEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fs, ok := request.([]*multipart.FileHeader)
		if !ok {
			return nil, ErrFileHeaderInvalid
		}
		return e(ctx, fs...)
	}
}
