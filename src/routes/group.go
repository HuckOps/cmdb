package routes

import (
	"cmdb/src/controllers"
	"github.com/gin-gonic/gin"
)

func GroupRoutes(e *gin.RouterGroup) {
	GroupRoutesMap := e.Group("/group")
	{
		GroupRoutesMap.GET("", controllers.GetGroupsList)
		GroupRoutesMap.POST("", controllers.CreateGroup)
		GroupRoutesMap.PUT("/:id", controllers.UpdateGroup)
		GroupRoutesMap.DELETE("/:id", controllers.DeleteGroup)
	}
}

func ServiceRoutes(e *gin.RouterGroup)  {
	ServiceRoutesMap := e.Group("/service")
	{
		ServiceRoutesMap.GET("", controllers.GetServiceList)
		ServiceRoutesMap.POST("", controllers.CreateService)
		ServiceRoutesMap.PUT("/:id", controllers.UpdateService)
		ServiceRoutesMap.DELETE("/:id", controllers.DeleteService)
	}
}

func MachineRoutes(e *gin.RouterGroup)  {
	MachineRoutesMap := e.Group("/machine")
	{
		MachineRoutesMap.GET("", controllers.GetMachineList)
		MachineRoutesMap.POST("/create_by_ips", controllers.CreateMachineByIPs)
	}
}