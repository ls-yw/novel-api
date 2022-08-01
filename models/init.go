package models

import (
	"novel/woodlsy/db"
)

type Model struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatedBy uint   `json:"created_by"`
	UpdatedBy uint   `json:"updated_by"`
}

var Orm db.Orm

//
// GetOne
// @Description: 获取单条记录
// @param m
// @param where
// @param orderBy
// @return int64
// @return error
//
func GetOne(m interface{}, where map[string]interface{}, orderBy string) (int64, error) {
	o := Orm.Model(m)
	if len(where) > 0 {
		o = o.Where(where)
	}
	result := o.Order(orderBy).Find(&m)
	return result.RowsAffected, result.Error
}

//
// GetList
// @Description: 获取分页列表
// @param m
// @param where
// @param orderBy
// @param offset
// @param limit
// @return int64
// @return error
//
func GetList(m interface{}, where map[string]interface{}, orderBy string, offset int, limit int) (int64, error) {
	o := Orm.Model(m)
	if len(where) > 0 {
		o = o.Where(where)
	}
	result := o.Order(orderBy).Offset(offset).Limit(limit).Find(m)
	return result.RowsAffected, result.Error
}

//
// GetAll
// @Description: 获取所有记录
// @param m
// @param where
// @param orderBy
// @return int64
// @return error
//
func GetAll(m interface{}, where map[string]interface{}, orderBy string) (int64, error) {
	o := Orm.Model(m)
	if len(where) > 0 {
		o = o.Where(where)
	}
	result := o.Order(orderBy).Find(m)
	return result.RowsAffected, result.Error
}

//
// GetCount
// @Description: 获取记录条数
// @param m
// @param where
// @return int64
//
func GetCount(m interface{}, where map[string]interface{}) int64 {
	var count int64
	o := Orm.Model(m)
	if len(where) > 0 {
		o = o.Where(where)
	}
	o.Count(&count)
	return count
}
