package server

import (
	"github.com/gin-gonic/gin"
	"kubernetes_management_system/pkg/server/middleware"
)

var (
	addr = "192.168.31.100:8080"
)

func InitServer() {
	r := gin.Default()

	r.Use(middleware.Cores())
	r.Use(gin.LoggerWithFormatter(middleware.LogsFormatDefine))

	group := r.Group("v1")
	{

	}

	r.Run(addr)
}
