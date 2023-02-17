package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"kubernetes_management_system/pkg/server"
	"kubernetes_management_system/pkg/signals"
	"os"
)

var (
	log = "gin.log"
)

func main() {

	_ = signals.SetupSignalHandler()

	//delete gin.log
	_ = os.Remove(log)

	file, err := os.Create(log)
	if err != nil {
		panic(err.Error())
	}

	//write log to file and console
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	server.InitServer()
}
