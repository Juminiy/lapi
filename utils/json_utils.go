package utils
import (
	"github.com/valyala/fastjson"
)
const (
	OK 					= 200
	Redirect 			= 302
	NonAuth				= 401
	NotFound 			= 404
	RequestFailure		= 500
	InternalServerError = 505
)
var (
	dataFormat JsonFormat
	JsonParser fastjson.Parser
)
type JsonFormat struct {
	Code 	int
	Message string
	Data 	interface{}
}
func (json *JsonFormat) Ok(data ... interface{}) JsonFormat {
	return JsonFormat{
		Code:    OK,
		Message: "OK",
		Data:    data,
	}
}
func (json *JsonFormat) Redirect(data ... interface{}) JsonFormat {
	return JsonFormat{
		Code: 	 Redirect,
		Message: "Redirect",
		Data: 	 data,
	}
}
func (json *JsonFormat) NonAuth(data ... interface{}) JsonFormat{
	return JsonFormat{
		Code: 	 NonAuth,
		Message: "NonAuth",
		Data:	 data,
	}
}
func (json *JsonFormat) None(data ... interface{}) JsonFormat {
	return JsonFormat{
		Code:    NotFound,
		Message: "NotFound",
		Data:    data,
	}
}
func (json *JsonFormat) RequestFail(data ... interface{}) JsonFormat {
	return JsonFormat{
		Code:    RequestFailure,
		Message: "RequestFailure",
		Data:    data,
	}
}
func (json *JsonFormat) Error(data ... interface{}) JsonFormat {
	return JsonFormat{
		Code: 	 InternalServerError,
		Message: "InternalServerError",
		Data: 	 data,
	}
}

