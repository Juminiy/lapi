package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"zhaoxin-api/config"
	"zhaoxin-api/context"
	"zhaoxin-api/mix_develop"
)

func RestApi(app *fiber.App){
	app.Get("/",context.OKContext)
	app.Get("/favicon.ico",proxy.Forward(config.Config("FAVICON_URL")))

	apiVersion_1 := app.Group("/v1")

	swagV1 := apiVersion_1.Group("/docs")
	SwagApi(swagV1)

	infoBase := apiVersion_1.Group("/info")
	InfoBaseApi(infoBase)

	authBase := apiVersion_1.Group("/auth")
	AuthBaseApi(authBase)

	utilsBase := apiVersion_1.Group("/utils")
	UtilsBaseApi(utilsBase)

	stuBase := apiVersion_1.Group("/stu")
	mix_develop.StuBaseApi(stuBase)

	// path notfound
	app.Use(context.NotFoundContext)
}