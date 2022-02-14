package main

import (
	"fmt"
	"service-man/common"
	"service-man/logman"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	config    *common.Configuration
	serverLog *logman.LogMan
	db        *gorm.DB
)

func init() {
	config = common.Config
	serverLog = common.ServerLog
	db = common.DB
}

func main() {
	application := iris.Default()
	app := RegisterRouter(application)
	app.Run(iris.Addr(fmt.Sprintf(":%d", config.HTTPPort)), iris.WithCharset("utf-8"))
}
