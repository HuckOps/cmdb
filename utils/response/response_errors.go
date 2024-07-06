package response

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type ResponseError interface {
	Handle (ctx *gin.Context)
	Panic () string
}

type Error struct {
	error error
}

func (e Error)Handle(ctx *gin.Context) {
	ctx.JSON(400, gin.H{
		"code": 400,
		"msg":  e.error.Error(),
	})
}

func (e Error) Panic() string {
	return e.error.Error()
}

func NewError(err string) ResponseError  {
	return Error{
		error: errors.New(err),
	}
}

var (
	HeaderError ResponseError = NewError("request header error")
)


