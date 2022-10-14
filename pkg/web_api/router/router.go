package router

import (
	"github.com/gin-gonic/gin"
)

func RegistryRouter (e *gin.Engine) {
	BaseRoute(e)
	ProjectRoute(e)
	LoginRoute(e)
	AdminRoute(e)
}
