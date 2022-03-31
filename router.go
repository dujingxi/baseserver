package main

import "github.com/kataras/iris/v12"

// Cors
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Method() == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization,File-Type,File-Name, Session-ID, Slice-Count, Slice-Number")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func RegisterRouter(app *iris.Application) *iris.Application {
	app.Use(Cors)
	common := app.Party("/")
	{
		common.Options("*", func(ctx iris.Context) {
			ctx.Next()
		})
	}
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
