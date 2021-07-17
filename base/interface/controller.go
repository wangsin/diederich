package baseinterface

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsin/diederich/base/consts"
	"net/http"
)

type BaseController interface {
	EchoJson(ctx *gin.Context, data interface{})
	EchoString(ctx *gin.Context, str string, data interface{})
	EchoJsonWithHttpCode(ctx *gin.Context, code int, data interface{})
	EchoJsonWithFormatStruct(ctx *gin.Context, bizCode int, data interface{})
	EchoJsonWithFormatStructAndHttpCode(ctx *gin.Context, httpCode, bizCode int, data interface{})
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

func (b BaseControllerClass) EchoJsonWithHttpCode(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
	ctx.Abort()
	return
}

func (b BaseControllerClass) EchoJsonWithFormatStruct(ctx *gin.Context, bizCode int, data interface{}) {
	trueData := make(map[string]interface{})
	trueData["msg"] = consts.CodeToName[bizCode]
	trueData["code"] = bizCode
	if data == nil {
		trueData["data"] = struct{}{}
	} else {
		trueData["data"] = data
	}
	b.EchoJson(ctx, trueData)
	return
}

func (b BaseControllerClass) EchoJsonWithFormatStructAndHttpCode(ctx *gin.Context, httpCode, bizCode int, data interface{}) {
	trueData := make(map[string]interface{})
	trueData["msg"] = consts.CodeToName[bizCode]
	trueData["code"] = bizCode
	if data == nil {
		trueData["data"] = struct{}{}
	} else {
		trueData["data"] = data
	}
	b.EchoJsonWithHttpCode(ctx, httpCode, trueData)
	return
}
