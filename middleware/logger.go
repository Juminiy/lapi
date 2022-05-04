package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func FileLogger(app *fiber.App) {
	// log file write into local
	file,err := os.OpenFile("./relative_file/api_request.log",os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err != nil {
		log.Fatalf("error opening file:%v",err )
	}
	app.Use(logger.New(logger.Config{
		Format:     "${time}|${pid}|${status} - ${method}|${path}\n",
		TimeFormat: "01-Jan-2001",
		TimeZone:   "Asia/Shanghai",
		Output: file,
	}))
}
func ConsoleLogger(app *fiber.App) {
	app.Use(logger.New())
}