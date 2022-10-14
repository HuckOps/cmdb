package server_init

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPermission (e gin.RoutesInfo){
	for _, route := range e{
		per_tmp := module.APIPermissionGroup{}
		result := db.MySQL.Model(&module.APIPermissionGroup{}).Where("api = ? and method = ?", route.Path, route.Method).First(&per_tmp)
		if result.Error == gorm.ErrRecordNotFound{
			permission := module.APIPermissionGroup{
				API: route.Path,
				Method: route.Method,
			}
			db.MySQL.Create(&permission)
		}
	}
}
