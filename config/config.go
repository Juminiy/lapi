package config

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"zhaoxin-api/utils"
)

var ServerConfigs = fiber.Config{
	Prefork: true,
	ServerHeader: Config("APP_NAME"),
	StrictRouting: true,
	CaseSensitive: true,
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return utils.ErrorResponse(ctx,err)
	},
	AppName: Config("APP_NAME"),
}

func Config(key string) string {
	return os.Getenv(key)
}
