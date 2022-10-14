package router

import (
	"cmdb/pkg/web_api/project"
	"cmdb/pkg/web_api/user"
	"github.com/gin-gonic/gin"
)

func ProjectRoute(e *gin.Engine) {
	projectRouteMap := e.Group("/api/v2/:project/project")
	projectRouteMap.Use(user.Auth())
	{
		/*
		路由组1： 项目群组路由
		*/
		projectRouteMap.GET("/group", user.AuthSuperAdminOrOthers(),project.GroupList)
		projectRouteMap.GET("/group/:id/permission", user.AuthSuperAdminOrOthers(), project.GroupPermission)
		projectRouteMap.GET("/group/:id", user.AuthSuperAdminOrOthers(),project.ProjectGroup)
		projectRouteMap.POST("/group", user.AuthSuperAdminOrOthers(), project.CreateGroup)


		/*
		路由组2： 项目相关信息
		*/
		projectRouteMap.GET("", user.AuthSuperAdminOrOthers(), project.ProjectInfo)
	}
}
