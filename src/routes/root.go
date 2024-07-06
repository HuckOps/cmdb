package routes

import (
	_ "cmdb/docs"
	"cmdb/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

var swagHandler gin.HandlerFunc

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}


func RootRoutes(e *gin.Engine) {
	//e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Group := e.Group("/api/v1")
	v1Group.Use(Cors())
	middleware.RegistryCommonMiddleware(v1Group)
	GroupRoutes(v1Group)
	ServiceRoutes(v1Group)
	MachineRoutes(v1Group)
	IDCRoutes(v1Group)
}
