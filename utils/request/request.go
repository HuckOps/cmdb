package request

import (
	"cmdb/utils/response"
	"github.com/gin-gonic/gin"
)

func GetProject(ctx *gin.Context) (err error) {
	project := ctx.GetHeader("AUTH_PROJECT")
	if project == "" {
		ctx.Abort()
		panic(response.HeaderError)
		return
	}
	ctx.Set("project", project)
	return
}

func GetProjectWithContext(ctx *gin.Context) (string) {
	project, exists := ctx.Get("project")
	if ! exists {
		panic(response.HeaderError)
	}
	return project.(string)
}
