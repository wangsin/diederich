package apptool

import (
	"github.com/gin-gonic/gin"
	"github.com/wangsin/diederich/app/tool/db"
	dbreq "github.com/wangsin/diederich/app/tool/db/req"
	"github.com/wangsin/diederich/base/consts"
	baseinterface "github.com/wangsin/diederich/base/interface"
)

type ToolController struct {
	baseinterface.BaseControllerClass
}

func (c *ToolController) EchoDBConf(ctx *gin.Context) {
	req := &dbreq.TestDBRequest{}
	err := ctx.Bind(req)
	if err != nil {
		c.EchoJsonWithFormatStruct(ctx, consts.ErrParams, err)
		return
	}

	c.EchoJsonWithFormatStruct(ctx, consts.Success, db.TestDB(req))
	return
}
