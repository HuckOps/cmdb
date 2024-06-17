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

type ResponseBody struct {
	Code ResponseCode `json:"code"`
	Data interface{}  `json:"data"`
	Msg  string       `json:"msg,omitempty"`
}

const BadRequestCode int = 400

var BadRequest = ResponseBody{
	Code: Failed,
	Data: map[string]string{},
	Msg:  "Bad request, pls check params and body",
}

var SearchError = ResponseBody{
	Code: Failed,
	Data: map[string]string{},
}

func Response(ctx *gin.Context, code int, body ResponseBody) {
	ctx.JSON(code, body)
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
	pagination.Limit = limit
	return
}
