package db

import (
	"gorm.io/gorm"
)

type Orm struct {
}

type OrmInterface interface {
	GetOne(interface{}, map[string]interface{}, string)
	GetList(interface{}, map[string]interface{}, string, int, int)
	GetAll(interface{}, map[string]interface{}, string)
	Count()
}

var db *gorm.DB

func OrmInit() {
	db = connect()
}

func (Orm) Model(value interface{}) *gorm.DB {
	return db.Model(value)
}

func (Orm) Insert(value interface{}) *gorm.DB {
	return db.Create(value)
}

func (Orm) Update(value interface{}, data map[string]interface{}, where map[string]interface{}) int64 {
	result := db.Model(value).Where(where).Updates(data)
	return result.RowsAffected
}
func (Orm) Deleted(value interface{}, where map[string]interface{}) int64 {
	result := db.Where(where).Delete(value)
	return result.RowsAffected
}
