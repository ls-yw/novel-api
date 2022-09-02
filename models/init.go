package models

import (
	"gorm.io/gorm"
	"novel/utils/common"
	"novel/woodlsy/db"
	"novel/woodlsy/log"
	"reflect"
	"strings"
)

type Model struct {
	Id        uint   `gorm:"primarykey" json:"id,omitempty"`
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
		whereSqlString, whereValueArray := parseWhere(where)
		if whereSqlString != "" {
			o.Where(whereSqlString, whereValueArray...)
		}
		//o = parseWhere(o, where)
		//o.Where(where)
	}
	if orderBy != "" {
		o.Order(orderBy)
	}
	if limit != 0 {
		o.Offset(offset).Limit(limit)
	}
	if fields != "" {
		o.Select(fields)
	}
	return o
}

func update(m interface{}, data map[string]interface{}, where map[string]interface{}) int64 {
	return Orm.Update(m, data, where)
}

func insert(m interface{}) {
	result := Orm.Insert(m)
	result.Debug()
	if result.Error != nil || result.RowsAffected == 0 {
		log.Logger.Error("新增记录失败", m, result.Error)
	}
}

//
// parseWhere
// @Description:解析where条件
//	map[string]interface{}{
//	"id":        2,                                                    // id =2
//	"cid":       []interface{}{"in", []int{1, 3}},                                         // cid in (1,2)
//	"pid":       []interface{}{"!=", 3},                               // pid != 3
//	"name":      []interface{}{"like", "%%大三%%"},                      // name like '%大三%'
//	"create_at": []interface{}{"between", "2022-08-01", "2022-08-02"}, // create_at between '2022-08-01' and '2022-08-02'
//	"bid":       []interface{}{"not in", []int{1, 2}},                 // bid not in (1, 2)
//	"or": map[string]interface{}{ // or链接
//	"id":        1,                                                    // id =2
//	"cid":       []interface{}{"in", []int{1, 3}},                                         // cid in (1,2)
//	"pid":       []interface{}{"!=", 4},                               // pid != 3
//	"name":      []interface{}{"like", "%%大三2%%"},                     // name like '%大三%'
//	"create_at": []interface{}{"between", "2022-08-03", "2022-08-04"}, // create_at between '2022-08-01' and '2022-08-02'
//	"bid":       []interface{}{"not in", []int{1, 3}},                 // bid not in (1, 2)
//	},
//	}
// @param where
// @return string
// @return []interface{}
//
func parseWhere(where map[string]interface{}) (string, []interface{}) {
	var whereSql []string
	var whereValue []interface{}
	var childWhereValue []interface{}
	var orWhereSqlString string
	for key, value := range where {
		if key == "or" && reflect.TypeOf(value).Kind() == reflect.Map {
			orWhereSqlString, childWhereValue = parseWhere(value.(map[string]interface{}))
		} else {

			switch vv := value.(type) {
			case []interface{}:
				switch strings.ToLower(vv[0].(string)) {
				case "!=", ">", ">=", "<", "<=", "like", "in", "not in":
					whereSql = append(whereSql, common.Join(" ", key, vv[0].(string), "?"))
					whereValue = append(whereValue, vv[1])
				case "between":
					whereSql = append(whereSql, common.Join(" ", key, vv[0].(string), "?", "and", "?"))
					whereValue = append(whereValue, vv[1], vv[2])
				}
			default:
				whereSql = append(whereSql, common.Join(" ", key, "=", "?"))

				whereValue = append(whereValue, value)
			}
		}
	}
	whereSqlString := strings.Join(whereSql, " AND ")

	var tmpSqlArray []string
	if len(whereSqlString) > 0 {
		tmpSqlArray = append(tmpSqlArray, common.Join(" ", "(", whereSqlString, ")"))
	}
	if len(orWhereSqlString) > 0 {
		tmpSqlArray = append(tmpSqlArray, common.Join(" ", "(", orWhereSqlString, ")"))
		return strings.Join(tmpSqlArray, " OR "), append(whereValue, childWhereValue...)
	}
	return whereSqlString, whereValue
}
