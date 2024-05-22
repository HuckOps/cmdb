package routes

import "github.com/gin-gonic/gin"

func GroupRoutes(e *gin.Engine) {
	GroupRoutesMap := e.Group("/api/v1/group")
	{
		GroupRoutesMap.Group("/")
	}
}
