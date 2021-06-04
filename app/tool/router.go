package apptool

import (
	"github.com/gin-gonic/gin"
)

type ToolRouterGroup struct {
	BasePath     string
	HandlerGroup []gin.HandlerFunc
}

func (t *ToolRouterGroup) RegRouter(engine *gin.Engine) {
	group := engine.Group(t.BasePath)
	for _, handlerFunc := range t.GetHandlers() {
		group.Use(handlerFunc)
	}

	group.GET("/dbConf", new(ToolController).EchoDBConf)
}

func (t *ToolRouterGroup) GetHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		func(context *gin.Context) {
			return
		},
	}
}
