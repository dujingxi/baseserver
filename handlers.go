/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:24:51
 * @Descripttion:
 */
package main

import (
	"github.com/kataras/iris/v12"
)

func TestHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"retCode": 200})
}

func PostServiceHandler(ctx iris.Context) {
	var m map[string]interface{}
	ctx.ReadJSON(&m)
	ctx.WriteString("post")
}

func VersionHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"code": 0, "errMsg": "success", "version": VERSION})
}
