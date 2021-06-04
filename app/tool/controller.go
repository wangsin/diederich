package apptool

import (
	"github.com/gin-gonic/gin"
	baseinterface "github.com/wangsin/diederich/base/interface"
)

type ToolController struct {
	baseinterface.BaseController
}

func (c *ToolController) EchoDBConf(ctx *gin.Context) {

}
