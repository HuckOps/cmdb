package model

import (
	"cmdb/pkg/mysql"
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
				if tx.Error != nil {
					fmt.Println(tx.Error.Error())
				}
			}
		}
	}
}

type Pagination struct {
	Skip  int `json:"skip,omitempty" default:"-1"`
	Limit int `json:"limit,omitempty" default:"10"`
}

type PaginationResult struct {
	Count int64 `json:"count"`
	Data  interface{}
}

func (p *Pagination) QueryPagination(model interface{}) (dest PaginationResult, err error) {
	tx := mysql.GormClient.Count(&dest.Count)
	if p.Limit >= 0 {
		tx = tx.Offset(p.Skip).Limit(p.Limit)
	}
	err = tx.Find(&dest.Data).Error
	return
}
