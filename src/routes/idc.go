package routes

import (
	"cmdb/src/controllers"
	"github.com/gin-gonic/gin"
)

func IDCRoutes(e *gin.RouterGroup){
//	数据中心
	IDCRoutesMap := e.Group("/idc")
	{
		IDCRoutesMap.GET("", controllers.GetIDCList)
		IDCRoutesMap.POST("", controllers.CreateIDC)
	}
}
