package models

import "novel/app/utils/common"

type BookApply struct {
	Model
	Uid      int    `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Author   string `json:"author,omitempty"`
	Platform string `json:"platform,omitempty"`
	Reply    string `json:"reply,omitempty"`
	BookId   int    `json:"book_id,omitempty"`
	ReplyAt  string `json:"reply_at,omitempty"`
	User     User   `gorm:"foreignKey:Uid"`
}

type ReturnApplyList struct {
	Name           string `json:"name,omitempty"`
	Author         string `json:"author,omitempty"`
	Reply          string `json:"reply,omitempty"`
	BookId         int    `json:"book_id,omitempty"`
	User__username string `json:"username"`
}

func (m BookApply) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []BookApply {
	list := make([]BookApply, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m BookApply) GetApplyList(offset int, size int) []ReturnApplyList {
	result := make([]ReturnApplyList, 0)
	Orm.Model(m).Joins("User").Order("id desc").Offset(offset).Limit(size).Find(&result)
	return result
}

func (m BookApply) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}

func (m BookApply) Insert() int {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
		m.ReplyAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}
