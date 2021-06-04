package baseinterface

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController interface {
	EchoJson(ctx *gin.Context, data interface{})
	EchoString(ctx *gin.Context, str string, data interface{})
}

type BaseControllerClass struct {
}

func (b BaseControllerClass) EchoJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
	return
}

func (b BaseControllerClass) EchoString(ctx *gin.Context, str string, data interface{}) {
	ctx.String(http.StatusOK, str, data)
	ctx.Abort()
	return
}
