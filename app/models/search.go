package models

import "novel/app/utils/common"

type Search struct {
	Model
	Keyword  string `json:"keyword,omitempty"`
	Platform string `json:"platform,omitempty"`
	Num      int    `json:"num,omitempty"`
}

func (m Search) GetOne(where map[string]interface{}, orderBy string, fields string) (info Search) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m Search) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []Search {
	list := make([]Search, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m Search) GetAll(where map[string]interface{}, orderBy string, fields string) []Search {
	list := make([]Search, 0)
	getAll(&list, where, orderBy, fields)
	return list
}

func (m Search) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}

func (m Search) Insert() int {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}

func (m Search) Update(data map[string]interface{}, where map[string]interface{}) int64 {
	data["updated_at"] = common.Now()
	return update(m, data, where)
}
