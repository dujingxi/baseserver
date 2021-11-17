package main

import (
	"baseserver/logman"
	"fmt"

	"github.com/kataras/iris/v12"
)

var (
	CONFIG *Configuration
	LOGMAN *logman.LogMan
)

func main() {
	application := iris.Default()
	app := RegisterRouter(application)
	app.Run(iris.Addr(fmt.Sprintf(":%d", CONFIG.HTTPPort)), iris.WithCharset("utf-8"))
}
