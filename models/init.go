package models

import (
	"gorm.io/gorm"
	"novel/woodlsy/db"
)

type Model struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	CreatedBy uint   `json:"created_by,omitempty"`
	UpdatedBy uint   `json:"updated_by,omitempty"`
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
func getOne(m interface{}, where map[string]interface{}, orderBy string, fields string) (int64, error) {
	o := sqlCondition(m, where, orderBy, 0, 0, fields)
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
func getList(m interface{}, where map[string]interface{}, orderBy string, offset int, limit int, fields string) (int64, error) {
	o := sqlCondition(m, where, orderBy, offset, limit, fields)
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
func getAll(m interface{}, where map[string]interface{}, orderBy string, fields string) (int64, error) {
	o := sqlCondition(m, where, orderBy, 0, 0, fields)
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
func getCount(m interface{}, where map[string]interface{}) int64 {
	var count int64
	o := sqlCondition(m, where, "", 0, 0, "")
	o.Count(&count)
	return count
}

//
// sqlCondition
// @Description: sql 条件组装
// @param m
// @param where
// @param orderBy
// @param offset
// @param limit
// @param fields
// @return *gorm.DB
//
func sqlCondition(m interface{}, where map[string]interface{}, orderBy string, offset int, limit int, fields string) *gorm.DB {
	o := Orm.Model(m)
	if len(where) > 0 {
		o = o.Where(where)
	}
	if orderBy != "" {
		o = o.Order(orderBy)
	}
	if limit != 0 {
		o = o.Offset(offset).Limit(limit)
	}
	if fields != "" {
		o = o.Select(fields)
	}
	return o
}

func Update() {

}

func Insert() {

}
