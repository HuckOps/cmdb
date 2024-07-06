package handle

import (
	"cmdb/pkg/restful"
	"cmdb/src/model"
	"cmdb/utils/request"
	"cmdb/utils/validate"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 合并两个函数的逻辑，通过使用opts参数来区分是否需要项目查询
func getResource[T any](ctx *gin.Context, params []string, useProject bool) {
	//result := &model.PaginationResult[T]{}
	//result.Data = dataModel

	pagination, err := restful.ParsePagination(ctx)
	if err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, restful.BadRequest)
		return
	}

	opts := []model.Option{}
	for _, param := range params {
		if v := ctx.Query(param); v != "" {
			opts = append(opts, model.SearchWithWhere(param+"="+v))
		}
	}

	// 如果需要项目查询
	if useProject {
		opts = append(opts, model.SearchWithProject(pagination.Project))
	}

	if result, err := model.SearchQueryPagination[T](pagination, opts...); err != nil {
		// 统一异常处理策略
		restful.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	} else{
		restful.Response(ctx, result)
	}
}

// 使用getResource进行资源获取，对于不需要项目的查询
func GetResourceWithoutProject[T any](ctx *gin.Context, params []string) {
	getResource[T](ctx, params, false)
}

// 使用getResource进行资源列表获取，对于需要项目的查询
func GetResourceList[T any](ctx *gin.Context, params []string) {
	getResource[T](ctx, params, true)
}

func CreateResource(ctx *gin.Context, data interface{})(err error){
	if err := validate.Validate(data); err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return err
	}
	if err := model.InsertRecord(data); err != nil {
		restful.ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return err
	}
	return nil
}

func UpdateResource(ctx *gin.Context, id string, data interface{})(error){
	idNum, err := strconv.Atoi(id)
	if  err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return err
	}

	if err := validate.Validate(data); err != nil {
		restful.ResponseError(ctx, http.StatusBadRequest, err.Error())
		return err
	}

	if rows := model.UpdateByID(data, uint(idNum), request.GetProjectWithContext(ctx)); rows == 0 {
		restful.ResponseError(ctx, http.StatusForbidden, "Forbidden")
		return errors.New("Forbidden")
	}
	return nil
}

func DeleteResource(ctx *gin.Context, id string, dataModel interface{}) (err error){
	idNum, err := strconv.Atoi(id)
	if err != nil {
		restful.ResponseError(ctx, 400, restful.BadRequest)
		return
	}
	if rows := model.DeleteByID(dataModel, uint(idNum), request.GetProjectWithContext(ctx)); rows == 0 {
		restful.ResponseError(ctx, http.StatusForbidden, "Forbidden")
		return errors.New("Forbidden")
	}
	return nil
}