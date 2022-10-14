package router

import (
	"cmdb/pkg/web_api/project"
	"cmdb/pkg/web_api/user"
	"github.com/gin-gonic/gin"
)

func AdminRoute(e *gin.Engine) {
	AdminRouteMap := e.Group("/api/v2/admin")
	AdminRouteMap.Use(user.Auth())
	{
		AdminRouteMap.POST("/project",user.Auth(), user.AuthSuperAdmin(),project.CreateProject)
		AdminRouteMap.POST("/user", user.Auth(), user.AuthSuperAdmin(), user.CreateUser)
		AdminRouteMap.POST("/permission/tag", user.Auth(), user.AuthSuperAdmin(), user.SetAPIPermissionTag)
	}
}
