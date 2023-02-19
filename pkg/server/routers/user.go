package routers

import (
	"github.com/gin-gonic/gin"
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
		//do nothing
		user.POST("/register", func(context *gin.Context) {
			context.String(500, "don't support")
		})
	}
}
