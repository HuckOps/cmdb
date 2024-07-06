package restful

import (
	"cmdb/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ResponseCode uint

const (
	Success = iota
	Failed
)

type ResponseBody[T []struct{}|struct{}|interface{}] struct {
	Code ResponseCode `json:"code"`
	Data T  `json:"data"`
	Msg  string       `json:"msg,omitempty"`
}

const BadRequestCode int = 400

const ServerError int = 500

var BadRequest = ResponseBody[map[string]string]{
	Code: Failed,
	Data: map[string]string{},
	Msg:  "Bad request, pls check params and body",
}

var Duplicate = ResponseBody[map[string]string]{
	Code: Failed,
	Data: map[string]string{},
	Msg:  "Duplicate key, pls check input data",
}

var DBOperateError = ResponseBody[interface{}]{
	Code: Failed,
	Data: map[string]string{},
	Msg:  "DB operate error",
}

func ErrorResponse(msg string) ResponseBody[interface{}]{
	return ResponseBody[interface{}]{
		Code: Failed,
		Data: map[string]string{},
		Msg:  msg,
	}
}

var SearchError = ResponseBody[interface{}]{
	Code: Failed,
	Data: map[string]string{},
}

func ResponseError(ctx *gin.Context, code int, body interface{}) {
	response := ResponseBody[interface{}]{
		Code: Failed,
		Data: body,
	}
	ctx.JSON(code, response)
}

func Response(ctx *gin.Context, body interface{}) {
	response := ResponseBody[interface{}]{
		Code: Success,
		Data: body,
	}
	ctx.JSON(200, response)
}

func ParsePagination(ctx *gin.Context) (pagination model.Pagination, err error) {
	skipStr := ctx.DefaultQuery("skip", "0")
	limitStr := ctx.DefaultQuery("limit", "-1")
	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		return pagination, err
	}
	pagination.Skip = skip
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return pagination, err
	}
	pagination.Project = ctx.GetHeader("AUTH_PROJECT")
	pagination.Limit = limit
	return
}
