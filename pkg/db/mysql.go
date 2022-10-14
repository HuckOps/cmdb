package db

import (
	"cmdb/pkg/config"
	"cmdb/pkg/module"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var MySQL *gorm.DB

func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.Username, config.Config.Mysql.Password, config.Config.Mysql.Host,
		config.Config.Mysql.Port, config.Config.Mysql.Db)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:      false,         // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("数据库建连失败")
	}
	MySQL = db
	db.AutoMigrate(&module.Project{}, &module.ProjectGroup{}, &module.User{}, &module.APIPermissionGroup{}, &module.User2Group{}, &module.GroupPermission{})
}
