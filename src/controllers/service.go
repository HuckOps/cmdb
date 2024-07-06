package controllers

import (
	"cmdb/pkg/restful"
	"cmdb/src/model"
	"cmdb/utils/handle"
	"cmdb/utils/request"
	"github.com/gin-gonic/gin"
)

func GetServiceList(ctx *gin.Context) {
	handle.GetResourceList[[]model.Service](ctx, []string{"name", "code", "rank"})
}

func CreateService(ctx *gin.Context)  {
	svc := &model.Service{}
	if err := ctx.ShouldBindJSON(svc); err != nil {
		restful.ResponseError(ctx, 400, restful.BadRequest)
		return
	}
	svc.Project = request.GetProjectWithContext(ctx)

	if err := handle.CreateResource(ctx, svc); err != nil {
		return
	}
	restful.Response(ctx, svc)
}

func UpdateService(ctx *gin.Context) {
	id := ctx.Param("id")
	svc := &model.Service{}
	if err := ctx.ShouldBindJSON(svc); err != nil {
		restful.ResponseError(ctx, 400, restful.BadRequest)
		return
	}
	svc.Project = request.GetProjectWithContext(ctx)
	if err := handle.UpdateResource(ctx, id, svc); err != nil {
		return
	}
	restful.Response(ctx, svc)
}

func DeleteService(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := handle.DeleteResource(ctx, id, &model.Service{}); err != nil {
		return
	}
	restful.Response(ctx, nil)
}