package mix_develop

import (
	"github.com/gofiber/fiber/v2"
	"zhaoxin-api/storage"
	"zhaoxin-api/utils"
)

type Stu struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
	Birth string `json:"birth"`
	Sex string `json:"sex"`
	Pic string `json:"pic"`
}

var (
	STU_HKEY = "STU"
)
func ListStu(amount int) {
	storage.Redis.HGetAll(STU_HKEY)
}
func AddStu(stu *Stu) error {
	//qResult,err := storage.Redis.Get(STU_HKEY+stu.Id ).Result()
	//fmt.Println("qResult",qResult,"err",err)
	//if err == nil && qResult == "" {
	//	sRes := storage.Redis.Set(STU_HKEY+stu.Id,stu,-1)
	//	return sRes.Err()
	//}
	storage.Redis.Set(STU_HKEY+stu.Id,stu,-1)
	return nil
}
func DelStu() {

}
func UpdateStu() {

}
func QStu() {

}

func ListStuContext(ctx *fiber.Ctx) error {
	return nil
}
func AddStuContext(ctx *fiber.Ctx) error {
	var stuInfo *Stu
	err := ctx.BodyParser(&stuInfo)
	if err != nil {
		return utils.ErrorResponse(ctx,"body parser",err)
	}
	err = AddStu(stuInfo)
	if err != nil {
		return utils.ErrorResponse(ctx,"add stu",err)
	}
	return utils.OkResponse(ctx,stuInfo,"添加成功")
}
func DelStuContext(ctx *fiber.Ctx) error {
	return nil
}
func UpdateStuContext(ctx *fiber.Ctx) error {
	return nil
}
func QStuContext(ctx *fiber.Ctx) error {
	return nil
}
func StuBaseApi(stu fiber.Router) {
	stu.Get("/",ListStuContext)
	stu.Post("/add",AddStuContext)
	stu.Post("/del",DelStuContext)
	stu.Post("/update",UpdateStuContext)
	stu.Get("/q",QStuContext)
}