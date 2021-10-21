package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
}

func NewContext(c *gin.Context) *Context {
	return &Context{c}
}

type sucResponse struct {
	Success bool `json:"success"`
}

func PackageResponse(res interface{}) (int, interface{}) {
	if res == nil {
		res = &sucResponse{
			Success: true,
		}
	}
	return http.StatusOK, res
}

func FailedResponse(c *Context, code Code, abort bool) {
	status := getHTTPStatus(code)
	body := gin.H{
		"code":    code,
		"message": code.String(),
	}
	if abort {
		c.AbortWithStatusJSON(status, body)
	} else {
		c.JSON(status, body)
	}
}

type HandlerFunc func(c *Context) (interface{}, error)

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		cc := NewContext(c)
		res, _ := handler(cc)
		c.PureJSON(PackageResponse(res))
	}
}
