package httpx

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChainDecoder(decode DecodeRequestFunc, others ...DecodeRequestFunc) DecodeRequestFunc {
	return func(c *gin.Context, req interface{}) error {
		var decodes []DecodeRequestFunc
		if decode != nil {
			decodes = append(decodes, decode)
		}
		decodes = append(decodes, others...)
		for _, decoder := range decodes {
			err := decoder(c, req)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func ShouldDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBind(req)
}

func JSONDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindJSON(req)
}

func XMLDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindXML(req)
}

func QueryDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindQuery(req)
}

func YAMLDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindYAML(req)
}

func HeaderDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindHeader(req)
}

func URIDecoder(c *gin.Context, req interface{}) error {
	return c.ShouldBindUri(req)
}

func JSONEncoder(c *gin.Context, response interface{}) error {
	c.JSON(http.StatusOK, response)
	return nil
}

func XMLEncoder(c *gin.Context, response interface{}) error {
	c.XML(http.StatusOK, response)
	return nil
}

func YAMLEncoder(c *gin.Context, response interface{}) error {
	c.YAML(http.StatusOK, response)
	return nil
}

func IndentedJSONEncoder(c *gin.Context, response interface{}) error {
	c.IndentedJSON(http.StatusOK, response)
	return nil
}

func SecureJSONEncoder(c *gin.Context, response interface{}) error {
	c.SecureJSON(http.StatusOK, response)
	return nil
}

func JSONPEncoder(c *gin.Context, response interface{}) error {
	c.SecureJSON(http.StatusOK, response)
	return nil
}

func AsciiJSONEncoder(c *gin.Context, response interface{}) error {
	c.AsciiJSON(http.StatusOK, response)
	return nil
}

func PureJSONEncoder(c *gin.Context, response interface{}) error {
	c.PureJSON(http.StatusOK, response)
	return nil
}

func ProtoBufEncoder(c *gin.Context, response interface{}) error {
	c.ProtoBuf(http.StatusOK, response)
	return nil
}
