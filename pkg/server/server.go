package server

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/middleware"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	r.Use(middleware.Cores())
	r.Use(gin.LoggerWithFormatter(middleware.LogsFormatDefine))

	publicGroup := r.Group("/api/v1")
	{
		routers.User(publicGroup)
	}

	common.LOG.Info("server addr: " + common.CONFIG.System.Addr)
	r.Run(common.CONFIG.System.Addr)
}
