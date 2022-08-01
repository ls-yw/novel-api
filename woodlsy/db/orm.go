package db

import "gorm.io/gorm"

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
