package baseinterface

import "github.com/gin-gonic/gin"

type Reg interface {
	RegRouter(*gin.Engine)
	GetHandlers() []gin.HandlerFunc
}
