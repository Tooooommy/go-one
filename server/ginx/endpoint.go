package ginx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"mime/multipart"
)

var ErrFileHeaderInvalid = errors.New("file header is invalid")

type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type FileEndpoint func(ctx context.Context, fs ...*multipart.FileHeader) (*JSONResponse, error)

type JSONEndpoint func(ctx context.Context, request interface{}) (*JSONResponse, error)

// JsonToEndpoint
func JsonToEndpoint(e JSONEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return e(ctx, request)
	}
}

// FileToEndpoint
func FileToEndpoint(e FileEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fs, ok := request.([]*multipart.FileHeader)
		if !ok {
			return nil, ErrFileHeaderInvalid
		}
		return e(ctx, fs...)
	}
}
