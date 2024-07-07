package routers

import (
	"github.com/gaochuang/cloudManagementSystem/handlers/system"
	"github.com/gin-gonic/gin"
)

func InitSystem(engine gin.Engine) {
	systemRouter := engine.Group("/api/v1/system")
	{
		systemRouter.POST("safe/settings", system.SetSystemSafe)
		systemRouter.GET("safe/settings", system.GetSystemSafe)
	}
}
