package model

import (
	"cmdb/pkg/mysql"
	"gorm.io/gorm"
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
		&Group{}, &Service{}, &IDC{}, &Room{}, &Cabinet{}, &Rack{}, &Device{}, &IP{}, &NetworkSegment{},
	}
	db.AutoMigrate(migrateList...)
}

type Pagination struct {
	Skip                int    `json:"skip,omitempty" default:"-1"`
	Limit               int    `json:"limit,omitempty" default:"10"`
	//EnableProjectFilter bool   `json:"-"`
	Project             string `json:"-"`
}

type PaginationResult[T interface{}] struct {
	Count int64       `json:"count"`
	Data  T `json:"data"`
}


type Options struct {
	DB *gorm.DB
}

type Option func(options *Options)

func NewOptions(tx *gorm.DB, opts ...Option) Options {
	options := Options{tx}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func SearchWithProject(t string) Option {
	return func(options *Options) {
		options.DB = options.DB.Where("project = ?", t)
	}
}

func SearchWithWhere(query interface{}, where ...interface{}) Option {
	return func(options *Options) {
		options.DB = options.DB.Where(query, where)
	}
}

func SearchQueryPagination[T any](pagination Pagination, opts ...Option) (*PaginationResult[T], error){
	var data T
	var total int64
	tx := mysql.GormClient.Model(data)
	//for _, w := range where {
	//	tx = tx.Where(w)
	//}
	//o := NewOptions(tx, SearchWithProject(pagination.Project))
	o := NewOptions(tx, opts...)
	err := o.DB.Count(&total).Error
	if err != nil {
		return nil, err
	}
	if pagination.Limit >= 0 {
		o.DB.Offset(pagination.Skip).Limit(pagination.Limit)
	}
	err = o.DB.Find(&data).Error
	result := &PaginationResult[T]{
		Data:  data,
		Count: total,
	}
	return result, err
}

func InsertRecord(data interface{}) error{
	tx := mysql.GormClient.Create(data).Scan(data)
	return tx.Error
}

func UpdateByID(data interface{}, id uint, project string) int64{
	// 更新记录并把新数据写入源结构体

	tx := mysql.GormClient.Model(data).Where("id = ? and project = ?", id, project).Updates(data).Scan(data)
	return tx.RowsAffected
}

func DeleteByID(model interface{}, id uint, project string) int64{
	tx := mysql.GormClient.Where("id = ? and project = ?", id, project).Delete(model)
	return tx.RowsAffected
}