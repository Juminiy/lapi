package router

import (
	"github.com/gofiber/fiber/v2"
	"lapi/context"
)

func UtilsBaseApi(utilsBase fiber.Router) {
	emailUtils := utilsBase.Group("/email")
	emailUtils.Post("/single",context.SendEmailContext)
	emailUtils.Post("/group",context.SendEmailGroup).Use(context.BasicAuth)

	ossUtils := utilsBase.Group("/oss")
	ossUtils.Get("/info",context.OSSInfo)
	ossUtils.Get("/dl",context.OSSDownloadFileDirect)
	ossUtils.Post("/ul",context.OSSUploadFileDirect)

}