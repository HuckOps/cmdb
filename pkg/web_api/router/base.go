package router

import (
	"cmdb/pkg/web_api/project"
	"cmdb/pkg/web_api/user"
	"github.com/gin-gonic/gin"
)

func BaseRoute(e *gin.Engine) {
	BaseRouteMap := e.Group("/api/v2")
	BaseRouteMap.Use(user.Auth())
	{
		BaseRouteMap.GET("/project", project.ProjectList)
	}
}

