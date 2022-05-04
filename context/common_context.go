package context

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofiber/fiber/v2"
	"lapi/config"
	"lapi/middleware"
	"lapi/model"
	"lapi/service"
	"lapi/storage"
	"lapi/utils"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)


func OKContext(ctx *fiber.Ctx) error {
	return utils.OkResponse(ctx,"Created by 2084Dev Team")
}
func NotFoundContext(ctx *fiber.Ctx) error {
	return utils.NoneResponse(ctx,"啥都木有")
}

func ApiInfoContext(ctx *fiber.Ctx) error {
	return utils.OkResponse(ctx,fiber.Map{
		"goVersion":         runtime.Version(),
		"description":     	 config.Config("APP_NAME"),
		"author":           "Chisato",
		"time":				"2022.3.10",
	})
}

func JWTAuth(ctx *fiber.Ctx) error {
	headerAuth := ctx.Get("Authorization")

	if headerAuth == "" {
		return utils.NonAuthResponse(ctx,"Need Header Authorization ")
	}
	chunks := strings.Split(headerAuth," ")
	if len(chunks) < 2 {
		return utils.NonAuthResponse(ctx,"Header Authorization Lacks")
	}
	user, err := middleware.Verify(chunks[1])
	if err != nil {
		return utils.NonAuthResponse(ctx,"Error occurs")
	}
	ctx.Locals("USER",user.ID)
	return ctx.Next()
}
func SendEmailContext(ctx *fiber.Ctx) error {
	emailDto := new(model.EmailSingleDto)
	if err := ctx.BodyParser(emailDto) ; err != nil {
		return utils.ErrorResponse(ctx,err)
	}
	if err := service.SendEmail(emailDto.To,emailDto.Subject,emailDto.Content,emailDto.Type) ; err != nil {
		return utils.ErrorResponse(ctx,"Send email failure",err)
	} else {
		return utils.OkResponse(ctx,"Send email success")
	}
}
func SendEmailGroup(ctx *fiber.Ctx) error {
	return utils.OkResponse(ctx)
}
func OSSInfo(ctx *fiber.Ctx) error{
	return utils.OkResponse(ctx,oss.Version)
}
func OSSUploadFileDirect(ctx *fiber.Ctx) error {
	bucket,err := storage.OSSClient.Bucket(config.Config("OSS_BUCKET"))
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	}
	file,err := ctx.FormFile("file")
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	}
	fileByte,err :=utils.FileParsing(file)
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	}
	filePath := filepath.Join("uploads",time.Now().In(config.CSTZone).Format("2001-01-01")) + "/" + file.Filename
	err = bucket.PutObject(filePath,bytes.NewReader(fileByte))
	if err != nil {
		return utils.ErrorResponse(ctx,err)
	} else {
		return utils.OkResponse(ctx,filePath)
	}
}
func OSSDownloadFileDirect(ctx *fiber.Ctx) error {
	_, err := storage.OSSClient.Bucket(config.Config("OSS_BUCKET"))
	if err != nil {
		return utils.ErrorResponse(ctx, err)
	}
	bucketFile := new(model.BucketFile)
	err = ctx.QueryParser(bucketFile)
	if err != nil {
		return utils.ErrorResponse(ctx, err)
	} else {
		return utils.OkResponse(ctx, "https://"+config.Config("OSS_BUCKET")+"."+config.Config("OSS_ENDPOINT")+"/"+bucketFile.ObjectFileName)
	}
}

