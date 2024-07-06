package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GID     uint   `gorm:"column:gid;not null;uniqueIndex:gid_project" union:"project" json:"gid" validate:"required"`
	Name    string `gorm:"column:name;type:varchar(50);not null" json:"name"  validate:"required"`
	Code    string `gorm:"column:code;type:varchar(50); not null;uniqueIndex:code_project" union:"project" json:"code" validate:"required"`
	Remark  string `gorm:"column:remark;type:text" json:"remark"`
	Project string `gorm:"column:project;type:varchar(50);not null;uniqueIndex:code_project,gid_project" json:"-"`
}

func (*Group) TableName() string {
	return "groups"
}
