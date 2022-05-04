package router


import (
	"github.com/gofiber/fiber/v2"
	"lapi/context"
)

func InfoBaseApi(infoBase fiber.Router) {
	infoBase.Get("/",context.ApiInfoContext)
}
