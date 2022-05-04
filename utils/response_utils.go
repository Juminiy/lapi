package utils

import "github.com/gofiber/fiber/v2"

func OkResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.Ok(data))
}
func RedirectResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.Redirect(data))
}
func NonAuthResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.NonAuth(data))
}
func NoneResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.None(data))
}
func RequestFailureResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.RequestFail(data))
}
func ErrorResponse(ctx *fiber.Ctx,data ... interface{}) error {
	return ctx.JSON(dataFormat.Error(data))
}
