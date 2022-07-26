package driver

import "gorm.io/gorm"

type Orm struct {
}

type OrmInterface interface {
	GetOne(interface{}, map[string]interface{}, string)
	GetList(interface{}, map[string]interface{}, string, int, int)
	GetAll(interface{}, map[string]interface{}, string)
}

var db *gorm.DB

func OrmInit() {
	db = connect()
}

func (Orm) GetOne(m interface{}, where map[string]interface{}, orderBy string) (int64, error) {
	result := db.Where(where).Order(orderBy).First(&m)
	return result.RowsAffected, result.Error
}

func (Orm) GetList(m []interface{}, where map[string]interface{}, orderBy string, offset int, row int) (int64, error) {
	result := db.Where(where).Order(orderBy).Offset(offset).Limit(row).Find(&m)
	return result.RowsAffected, result.Error
}

func (Orm) GetAll(m []interface{}, where map[string]interface{}, orderBy string) (int64, error) {
	result := db.Where(where).Order(orderBy).Find(&m)
	return result.RowsAffected, result.Error
}
