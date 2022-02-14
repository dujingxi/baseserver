package main

import "github.com/kataras/iris/v12"

func TestHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"retCode": 200})
}

func PostServiceHandler(ctx iris.Context) {

}

func GetServiceHandler(ctx iris.Context) {

}
func PutServiceHandler(ctx iris.Context) {

}
func DelServiceHandler(ctx iris.Context) {

}
