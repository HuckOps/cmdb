package user

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"github.com/gin-gonic/gin"
)

type SetAPIPermissionTagRequest struct {
	API string `json:"api"`
	Method string `json:"method"`
	Tag string `json:"tag"`
}

func SetAPIPermissionTag(c *gin.Context)  {
	request := SetAPIPermissionTagRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"code": 1, "msg": "Request body error", "data": []string{}})
		return
	}
	db.MySQL.Model(&module.APIPermissionGroup{}).Where("api = ? and method =?", request.API, request.Method).
		Update("tag", request.Tag)
	c.JSON(200, gin.H{"code": 0, "msg": "Set tag success", "data": []string{}})
}
