package models

import "novel/app/utils/common"

type User struct {
	Model
	Mobile   string `json:"mobile,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
	LastIp   string `json:"last_ip,omitempty"`
	LastTime string `json:"last_time,omitempty"`
	Count    int16  `json:"count,omitempty"`
	IsVip    int8   `json:"is_vip,omitempty"`
}

func (m User) GetOne(where map[string]interface{}, orderBy string, fields string) (info User) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m User) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []User {
	list := make([]User, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m User) GetAll(where map[string]interface{}, orderBy string, fields string) []User {
	list := make([]User, 0)
	getAll(&list, where, orderBy, fields)
	return list
}

func (m User) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}

func (m User) Insert() int {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}

func (m User) Update(data map[string]interface{}, where map[string]interface{}) int64 {
	data["updated_at"] = common.Now()
	return update(m, data, where)
}
