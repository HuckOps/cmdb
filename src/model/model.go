package model

import (
	"cmdb/src/config"
	"cmdb/utils"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

func parseGormTag(tag string) map[string]string {
	tags := make(map[string]string)
	pairs := strings.Split(tag, ";")
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			tags[kv[0]] = kv[1]
		} else {
			tags[kv[0]] = ""
		}
	}
	return tags
}

func Migrate(db *gorm.DB) {
	migrateList := []interface{}{
		&Group{},
	}
	db.AutoMigrate(migrateList...)
	for _, model := range migrateList {
		m := reflect.TypeOf(model).Elem()
		for i := 0; i < m.NumField(); i++ {
			field := m.Field(i)
			unionTags := field.Tag.Get("union")
			UnionTagsList := strings.Split(unionTags, ",")

			gormTag := field.Tag.Get("gorm")
			columnName := field.Name
			tableName := model.(interface {
				TableName() string
			}).TableName()
			if gormTag != "" {
				tags := parseGormTag(gormTag)
				if col, ok := tags["column"]; ok {
					columnName = col
				}
			}

			if config.CmdArgs.Project && utils.Contains(UnionTagsList, "project") {
				sql := fmt.Sprintf("CREATE INDEX %s ON `%s`(%s, project)", fmt.Sprintf("%s_project", columnName), tableName, columnName)
				tx := db.Exec(sql)
				fmt.Println(tx.Error.Error())
			}
		}
	}
}
