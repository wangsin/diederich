package mgin

import (
	"github.com/gin-gonic/gin"
	apptool "github.com/wangsin/diederich/app/tool"
	baseinterface "github.com/wangsin/diederich/base/interface"
)

var GinEngine *gin.Engine

func Init() error {
	GinEngine = gin.Default()

	regList := []baseinterface.Reg{
		&apptool.ToolRouterGroup{
			BasePath: "/diederich/tool/",
		},
	}

	for _, reg := range regList {
		reg.RegRouter(GinEngine)
	}

	return nil
}
