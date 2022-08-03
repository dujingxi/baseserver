/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:19:21
 * @Descripttion:
 */
package main

import (
	"baseserver/common"
	"baseserver/logman"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

var (
	settingConfig *common.Configuration
	fileConfig    *common.Configuration
	db            *gorm.DB
	serverLog     *logman.LogMan
)

func init() {
	cf := flag.String("f", "conf.json", "specify the config file.")
	flag.Parse()
	settingConfig = HandleConfig(*cf)

	// Initialize the mysql db
	db = common.InitDB(settingConfig)
	// db = common.DB

	// for a log file
	serverLog = logman.NewLogMan(filepath.Join(settingConfig.LogDir, "server.log"))
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
	app.Run(iris.Addr(fmt.Sprintf("%v:%v", settingConfig.HTTPBind, settingConfig.HTTPPort)), iris.WithCharset("utf-8"))
}
