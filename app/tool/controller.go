package apptool

import (
	"github.com/gin-gonic/gin"
	dbreq "github.com/wangsin/diederich/app/tool/db/req"
	baseinterface "github.com/wangsin/diederich/base/interface"
)

type ToolController struct {
	baseinterface.BaseController
}

func (c *ToolController) EchoDBConf(ctx *gin.Context) {
	req := &dbreq.TestDBRequest{}
	err := ctx.BindJSON(req)
	if err != nil {
		c.EchoJson(ctx, map[string]string{
			"code": "40000001",
			"msg":  "参数错误",
			"data": "",
		})
		return
	}

}
