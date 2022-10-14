package server_init

import (
	"cmdb/pkg/db"
	"cmdb/pkg/module"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func Init () {
	sa_project := module.Project{Code: "cmdb", Name: "运维配置平台"}
	if err := db.MySQL.Where(&sa_project).First(&sa_project).Error; err == gorm.ErrRecordNotFound{
		db.MySQL.Create(&sa_project)
		log.Println("cmdb项目已创建")
		//db.MySQL.Where(&sa_project).First(&sa_project)
	}
	sa_project_group := module.ProjectGroup{Name: "运维", Code: "cmdb.sa", Project: sa_project.ID}
	if err := db.MySQL.Where(&sa_project_group).First(&sa_project_group).Error; err == gorm.ErrRecordNotFound{
		db.MySQL.Create(&sa_project_group)
		log.Println("cmdb运维组已创建")
	}

	user := module.User{Name: "admin", UserName: "admin"}
	if err := db.MySQL.Where(&user).First(&user).Error; err == gorm.ErrRecordNotFound{
		hasedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		user.Password = string(hasedPassword)
		db.MySQL.Create(&user)
		db.MySQL.Create(&module.User2Group{UserID: user.ID, GroupID: sa_project_group.ID})
	}

}
