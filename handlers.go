package main

import "github.com/kataras/iris/v12"

func TestHandler(ctx iris.Context) {
	ctx.JSON(iris.Map{"retCode": 200})
}
