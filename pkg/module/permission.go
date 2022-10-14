package module

import (
	"gorm.io/gorm"
)

type APIPermissionGroup struct {
	ID int `gorm:"column:id;primaryKey"`
	API string `gorm:"column:api;not null;type:varchar(100)"`
	Method string `gorm:"column:method;not null;type:varchar(20)"`
	Tag string `gorm:"column:tag;type:varchar(50)"`
	gorm.Model
}

func (APIPermissionGroup) TableName() string {
	return "api_permission_group"
}