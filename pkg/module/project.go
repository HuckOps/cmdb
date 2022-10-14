package module

import "gorm.io/gorm"

type Project struct {
	ID int `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"not null;type:varchar(40);column:name; unique" json:"name"`
	Code string `gorm:"not null;type:varchar(20);column:code; unique" json:"code"`
	gorm.Model
}
func (Project) TableName() string {
	return "project"
}

type ProjectGroup struct {
	ID int `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"not null;type:varchar(40);column:name;unique" json:"name"`
	Code string `gorm:"not null;type:varchar(20);column:code;unique" json:"code"`
	Project int `gorm:"not null"`
	gorm.Model
}

func (ProjectGroup) TableName() string {
	return "project_group"
}