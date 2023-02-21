package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"kubernetes_management_system/common"
	"kubernetes_management_system/pkg/server"
	"kubernetes_management_system/pkg/signals"
	"kubernetes_management_system/tool"
	"os"
)

var (
	log = "gin.log"
)

func main() {

	_ = signals.SetupSignalHandler()

	file := createGinLog()
	initConfig()
	//write log to file and console
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	server.InitServer()

	defer func() {
		db, _ := common.DB.DB()
		db.Close()
	}()
}

func initConfig() {
	common.VIPER = tool.Viper()
	common.LOG = tool.Zap()
	common.DB = tool.GormMysql()
}

func createGinLog() *os.File {
	//delete gin.log
	_ = os.Remove(log)
	file, err := os.Create(log)
	if err != nil {
		panic(err.Error())
	}
	return file
}
