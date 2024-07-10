package routers

import (
	"github.com/gaochuang/cloudManagementSystem/handlers/role"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(engine *gin.Engine) {
	roleRouter := engine.Group("/api/v1/role")
	{
		roleRouter.GET("", role.ListRoles)

	}
}
