package server

import (
	"github.com/gin-gonic/gin"
	"study-go/web/common"
)

func NoRouteHandler(c *gin.Context) {
	common.FailedResponse(common.NewContext(c), common.ErrRequestMethodNotFound, true)
}

func NoMethodHandler(c *gin.Context) {
	common.FailedResponse(common.NewContext(c), common.ErrRequestMethodNotFound, true)
}

func Health(c *gin.Context) {
	c.JSON(common.PackageResponse(nil))
}
