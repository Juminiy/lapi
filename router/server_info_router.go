package router


import (
"github.com/gofiber/fiber/v2"
"zhaoxin-api/context"
)

func InfoBaseApi(infoBase fiber.Router) {
	infoBase.Get("/",context.ApiInfoContext)
}
