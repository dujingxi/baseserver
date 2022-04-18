/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-04-18 14:36:06
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

func GetServiceHandler(ctx iris.Context) {

}
func PutServiceHandler(ctx iris.Context) {

}
func DelServiceHandler(ctx iris.Context) {

}
