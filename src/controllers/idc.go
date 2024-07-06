package controllers

import (
	"cmdb/pkg/restful"
	"cmdb/src/model"
	"cmdb/utils/handle"
	"github.com/gin-gonic/gin"
)

// GetIDCList
// @Summary      Get IDC List
// @Description  Get IDC List
// @Tags         idc
// @Accept       json
// @Produce      json
// @Param        code  query   string  false  "IDC Code"
// @Param        name  query   string  false  "IDC Name"
// @Success      200  {object}  restful.ResponseBody[model.PaginationResult[[]model.IDC]]
// @Router       /idc [get]
func GetIDCList(ctx *gin.Context){
	handle.GetResourceWithoutProject[[]model.IDC](ctx, []string{"code", "name", "vendor"} )
}

// CreateIDC
// @Summary      Create IDC
// @Description  Create IDC
// @Tags         idc
// @Accept       json
// @Produce      json
// @Param        idc  body   model.IDC  true  "IDC"
// @Success      200  {object}  restful.ResponseBody[model.IDC]
// @Router       /idc [post]
func CreateIDC(ctx *gin.Context) {
	idc := model.IDC{}
	if err := ctx.ShouldBindJSON(&idc); err != nil {
		restful.ResponseError(ctx, restful.BadRequestCode, err)
		return
	}
	if err := handle.CreateResource(ctx, &idc); err != nil {
		return
	}
	restful.Response(ctx, idc)
}
