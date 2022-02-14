package main

import "github.com/kataras/iris/v12"

func RegisterRouter(app *iris.Application) *iris.Application {
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
