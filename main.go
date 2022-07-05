/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-07-05 11:19:30
 * @Descripttion:
 */
package main

import (
	"fmt"
	"path/filepath"
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
	// serverLog = common.ServerLog
	db = common.DB

	// for a log file
	serverLog = logman.NewLogMan(filepath.Join(common.Config.LogDir, "server.log"))
	serverLog.SetSaveMode(logman.BySize)
	serverLog.SetSaveVal(20)
	// ServerLog.SetLevel(logman.DEBUG)
}

func main() {
	// serverLog.Fatalf(logman.Fields{
	// 	"message": "log fatal",
	// })
	// serverLog.Errorf(logman.Fields{
	// 	"message": "log error",
	// })
	// serverLog.Warnf(logman.Fields{
	// 	"message": "log warn",
	// })
	// serverLog.Infof(logman.Fields{
	// 	"message": "log info",
	// })
	// serverLog.Debugf(logman.Fields{
	// 	"message": "log debug",
	// })
	application := iris.Default()
	app := RegisterRouter(application)
	app.Run(iris.Addr(fmt.Sprintf(":%d", config.HTTPPort)), iris.WithCharset("utf-8"))
}
