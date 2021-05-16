package ginx

import (
	"context"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

// JSONResponse
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NoDecoder
func NoDecoder(c *gin.Context, request interface{}) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		return request, nil
	}
}

// NoEncoder
func NoEncoder(c *gin.Context) httptransport.EncodeResponseFunc {
	return func(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
		return nil
	}
}

// JSONDecoder
func JSONDecoder(c *gin.Context, request interface{}) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		if err := c.BindHeader(request); err != nil {
			return nil, err
		}
		if err := c.BindUri(request); err != nil {
			return nil, err
		}
		if err := c.BindQuery(request); err != nil {
			return nil, err
		}
		if err := c.ShouldBind(request); err != nil {
			return nil, err
		}
		return request, nil
	}
}

// JSONEncoder
func JSONEncoder(c *gin.Context) httptransport.EncodeResponseFunc {
	return func(ctx context.Context, writer http.ResponseWriter, response interface{}) (err error) {
		if response == nil {
			err = ErrReturnIsNil
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.JSON(http.StatusOK, response.(JSONResponse))
		}
		return err
	}
}

// FileDecoder
func FileDecoder(c *gin.Context, resp interface{}) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, req *http.Request) (interface{}, error) {
		fs, err := c.MultipartForm()
		if err != nil {
			return nil, err
		}
		return fs.File["file"], nil
	}
}
