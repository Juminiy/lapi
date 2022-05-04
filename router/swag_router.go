package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwagApi(swag fiber.Router) {
	swag.Get("/*",swagger.HandlerDefault)
}
