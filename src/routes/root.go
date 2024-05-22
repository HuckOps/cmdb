package routes

import "github.com/gin-gonic/gin"

func RootRoutes(e *gin.Engine) {
	GroupRoutes(e)
}
