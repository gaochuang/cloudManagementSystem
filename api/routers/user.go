package routers

import (
	"github.com/gaochuang/cloudManagementSystem/api/middleware"
	user2 "github.com/gaochuang/cloudManagementSystem/handlers/user"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitializePublicRoutes(engine *gin.Engine) {
	engine.Use(middleware.CoreMiddleware())
	guest := engine.Group("/api/v1/platform")
	{
		guest.GET("/ping", func(c *gin.Context) {
			c.String(200, "ping")
		})
		guest.GET("/address", func(c *gin.Context) {
			c.String(200, c.Request.RemoteAddr)
		})
	}

	login := engine.Group("/api/v1/user")
	{
		login.POST("/register", user2.Register)
		login.POST("/login", user2.Login)
	}

	//为prometheus提供metrics数据, 请求数量、响应时间、错误率等
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func InitUserRouter(engin *gin.Engine) {
	user := engin.Group("/api/v1/user")
	{
		user.GET("info", user2.UserInfo)
	}

}
