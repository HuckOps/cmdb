package controllers

import (
	"cmdb/pkg/restful"
	"cmdb/src/model"
	"github.com/gin-gonic/gin"
)

func GetGroupsList(ctx *gin.Context) {
	pagination, err := restful.ParsePagination(ctx)
	if err != nil {
		restful.Response(ctx, restful.BadRequestCode, restful.BadRequest)
		return
	}

	dest, err := pagination.QueryPagination([]model.Group{})
	if err != nil {
		return
	}
}
