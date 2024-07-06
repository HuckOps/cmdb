package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(50);not null;" json:"name" validate:"required"`
	Code string `gorm:"column:code;type:varchar(50);not null;uniqueIndex:code_project" json:"code" validate:"required"`
	Rank uint   `gorm:"column:rank;type:int(11);not null" json:"rank" validate:"min=0,max=5,required"`
	Remark string `gorm:"column:remark;type:text" json:"remark"`
	Project string `gorm:"column:project;type:varchar(50);not null;uniqueIndex:code_project" json:"-"`
}

func (*Service) TableName() string {
	return "services"
}