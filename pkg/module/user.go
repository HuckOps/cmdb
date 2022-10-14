package module

import (
	"gorm.io/gorm"
)

type User struct {
	ID int `gorm:"primaryKey"`
	UserName string `gorm:"column:username;not null;unique;type:varchar(40)"`
	Password string `gorm:"column:password;not null;type:varchar(255)"`
	Name string `gorm:"column:name;not null;type:varchar(50)"`
	FullName string `gorm:"column:fullname;type:varchar(100)"`
	gorm.Model
}
func (User) TableName() string {
	return "user"
}
type User2Group struct {
	ID int `gorm:"primaryKey"`
	UserID int `gorm:"not null"`
	GroupID int `gorm:"not null"`
}
func (User2Group) TableName() string {
	return "user2group"
}

type GroupPermission struct {
	ID int `gorm:"primaryKey"`
	GroupID int `gorm:"not null"`
	PermissionTag string `gorm:"not null"`
}

func (GroupPermission) TableName() string {
	return "group_permission"
}