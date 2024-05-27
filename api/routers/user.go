package routers

import (
	router "github.com/gaochuang/cloudManagementSystem/pkg/server/service/user"
	"github.com/gin-gonic/gin"
)

func InitializePublicRoutes(engine *gin.Engine) {
	guest := engine.Group("/api/v1/platform")
	{
		guest.GET("/ping", func(c *gin.Context) {
			c.String(200, "ping")
		})
		guest.GET("/address", func(c *gin.Context) {
			c.String(200, c.Request.RemoteAddr)
		})
	}
}

func User(group *gin.RouterGroup) {
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
