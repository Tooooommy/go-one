package ginx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NopDecoder(request interface{}) DecodeFunc {
	return func(context *gin.Context) (interface{}, error) {
		return request, nil
	}
}

func ShouldDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBind(request)
		return request, err
	}
}

func JSONDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindJSON(request)
		return request, err
	}
}

func XMLDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindXML(request)
		return request, err
	}
}

func QueryDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindQuery(request)
		return request, err
	}
}

func YAMLDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindYAML(request)
		return request, err
	}
}

func HeaderDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindHeader(request)
		return request, err
	}
}

func URIDecoder(request interface{}) DecodeFunc {
	return func(ctx *gin.Context) (interface{}, error) {
		err := ctx.ShouldBindUri(request)
		return request, err
	}
}

func NopEncoder(ctx *gin.Context, response interface{}) error {
	return nil
}

func JSONEncoder(ctx *gin.Context, response interface{}) error {
	ctx.JSON(http.StatusOK, response)
	return nil
}

func XMLEncoder(ctx *gin.Context, response interface{}) error {
	ctx.XML(http.StatusOK, response)
	return nil
}

func YAMLEncoder(ctx *gin.Context, response interface{}) error {
	ctx.YAML(http.StatusOK, response)
	return nil
}

func IndentedJSONEncoder(ctx *gin.Context, response interface{}) error {
	ctx.IndentedJSON(http.StatusOK, response)
	return nil
}

func SecureJSONEncoder(ctx *gin.Context, response interface{}) error {
	ctx.SecureJSON(http.StatusOK, response)
	return nil
}

func JSONPEncoder(ctx *gin.Context, response interface{}) error {
	ctx.SecureJSON(http.StatusOK, response)
	return nil
}

func AsciiJSONEncoder(ctx *gin.Context, response interface{}) error {
	ctx.AsciiJSON(http.StatusOK, response)
	return nil
}

func PureJSONEncoder(ctx *gin.Context, response interface{}) error {
	ctx.PureJSON(http.StatusOK, response)
	return nil
}

func ProtoBufEncoder(ctx *gin.Context, response interface{}) error {
	ctx.ProtoBuf(http.StatusOK, response)
	return nil
}
