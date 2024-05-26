/*
Copyright © 2024 Gao chuang <ienjoyarmlinux@163.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/gaochuang/cloudManagementSystem/common"
	"github.com/gaochuang/cloudManagementSystem/pkg/server/routers"
	"github.com/gaochuang/cloudManagementSystem/pkg/signals"
	"github.com/gaochuang/cloudManagementSystem/tool"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var (
	log = "gin.log"
)

func main() {
	_ = signals.SetupSignalHandler()
	file := createGinLog()
	initConfig()

	//init sql table
	common.MysqlTable(common.DB)
	//write log to file and console
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	routers.InitServer()

	defer func() {
		db, _ := common.DB.DB()
		db.Close()
	}()
}

func initConfig() {
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
