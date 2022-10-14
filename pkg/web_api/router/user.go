package router

import (
	"cmdb/pkg/web_api/user"
	"github.com/gin-gonic/gin"
)

func LoginRoute(e *gin.Engine){
	LoginRouteMap := e.Group("/api/v2/login")
	{
		LoginRouteMap.POST("", user.Login)
	}
}

