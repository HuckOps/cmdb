package controllers

import (
	"cmdb/pkg/logger"
	"cmdb/pkg/restful"
	"cmdb/src/model"
	"cmdb/utils/handle"
	"cmdb/utils/request"
	"github.com/gin-gonic/gin"
)

func GetGroupsList(ctx *gin.Context) {
	handle.GetResourceList[[]model.Group](ctx, []string{"name", "code", "gid"})
}

func CreateGroup(ctx *gin.Context) {
	group := model.Group{}
	if err := ctx.ShouldBindBodyWithJSON(&group); err != nil {
		logger.ServerLogger.Error(err.Error())
		restful.ResponseError(ctx, 400, restful.BadRequest)
		return
	}
	group.Project = request.GetProjectWithContext(ctx)

	if err := handle.CreateResource(ctx, &group); err != nil {
		return
	}
	restful.Response(ctx, group)
}

// UpdateGroup 更新群组
func UpdateGroup(ctx *gin.Context) {
	id := ctx.Param("id")
	group := &model.Group{}
	if err := ctx.ShouldBindBodyWithJSON(group); err != nil {
		logger.ServerLogger.Error(err.Error())
		restful.ResponseError(ctx, 400, restful.BadRequest)
		return
	}
	group.Project = request.GetProjectWithContext(ctx)
	if err := handle.UpdateResource(ctx, id, group); err != nil{
		return
	}
	restful.Response(ctx, group)
}

func DeleteGroup(ctx *gin.Context)  {
	id := ctx.Param("id")
	if err := handle.DeleteResource(ctx, id, &model.Group{}); err != nil{
		return
	}
	restful.Response(ctx, nil)
}