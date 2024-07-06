package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GormClient *gorm.DB

func InitMySQL() {
	dsn := "root:sjh080815@tcp(127.0.0.1:3306)/cmdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version

	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	//if err != nil {
	//	logger.ServerLogger.Fatal(fmt.Sprintf("can't connect to db: %v", err.Error()))
	//	os.Exit(1)
	//}
	//logger.ServerLogger.Info(fmt.Sprintf("Connect to db success, dsn: %s", dsn))
	GormClient = db
}
