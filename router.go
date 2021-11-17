package main

import "github.com/kataras/iris/v12"

func RegisterRouter(app *iris.Application) *iris.Application {
	api := app.Party("/api")
	{
		api.Get("/", TestHandler)
	}
	return app
}
