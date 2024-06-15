package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(engine *gin.Engine) {
	//gin.Recovery() 当程序中出现未知异常，返回http code 500
	engine.Use(gin.Recovery(), AuthMiddleware(), MetricsExportMiddleware(), CoreMiddleware())
}
