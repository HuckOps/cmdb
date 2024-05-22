package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GID     uint   `gorm:"column:gid;not null" union:"project"`
	Name    string `gorm:"column:name;type:varchar(50);not null"`
	Code    string `gorm:"column:code;type:varchar(50); not null;" union:"project"`
	Remark  string `gorm:"column:remark;type:text"`
	Project string `gorm:"column:project;type:varchar(50)"`
}

func (*Group) TableName() string {
	return "groups"
}
