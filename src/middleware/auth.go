package middleware

import (
	"cmdb/utils/request"
	"github.com/gin-gonic/gin"
)

func RegistryCommonMiddleware(e *gin.RouterGroup) {
	e.Use(gin.Recovery())
	e.Use(Auth())
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// project鉴权
		err := request.GetProject(context)
		if err != nil {
			context.Abort()
		}
	}
}
