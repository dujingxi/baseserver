/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:19:57
 * @Descripttion:
 */
package main

import (
	"baseserver/logman"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
)

// Cors
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Method() == http.MethodOptions {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization,File-Type,File-Name, Session-ID, Slice-Count, Slice-Number")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
func LogReqBody(ctx iris.Context) {
	requestLog := fmt.Sprintf("REQ %v %v %v", ctx.Request().RemoteAddr, ctx.Method(), ctx.Request().URL)
	if ctx.Method() == http.MethodPost || ctx.Method() == http.MethodPut || ctx.Method() == http.MethodPatch {
		params := ""
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = ioutil.NopCloser(buf)
			params = string(body)
			if strings.Contains(params, "\r\n") {
				params = strings.ReplaceAll(params, "\r\n", "")
			}
			if strings.Contains(params, "\n") {
				params = strings.ReplaceAll(params, "\n", "")
			}
			params = strings.ReplaceAll(params, " ", "")
			requestLog += " - " + params
		}

	}
	serverLog.Infof(logman.Fields{
		"message": requestLog,
	})

	ctx.Next()
}

func RegisterRouter(app *iris.Application) *iris.Application {
	if settingConfig.CrosConfig {
		app.Use(Cors)
		common := app.Party("/")
		{
			common.Options("*", func(ctx iris.Context) {
				ctx.Next()
			})
		}
	}
	app.Use(LogReqBody)
	api := app.Party("/api")
	{
		api.Get("/", TestHandler)
		service := api.Party("/service")
		{
			service.Post("/", PostServiceHandler)
			service.Get("/", GetServiceHandler)
			service.Get("/{id:int}", GetServiceHandler)
			service.Put("/{id:int}", PutServiceHandler)
			service.Delete("/{id:int}", DelServiceHandler)
		}
	}
	return app
}
