package routers

import (
	"github.com/gin-gonic/gin"
	router "kubernetes_management_system/pkg/server/service/user"
)

func User(group *gin.RouterGroup) {
	guest := group.Group("/wukong")
	{
		guest.GET("/ping", func(context *gin.Context) {
			context.String(200, "pong")
		})
		guest.GET("/addr", func(context *gin.Context) {
			context.String(200, context.Request.RemoteAddr)
		})
	}

	user := group.Group("/user")
	{
		user.POST("/register", router.Register)
		user.POST("/login", router.Login)
	}
}

func InitUserInfo(group *gin.RouterGroup) {
	routerGroup := group.Group("user")
	{
		routerGroup.GET("info", router.UserInfo)
	}

}
