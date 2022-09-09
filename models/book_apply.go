package models

import "novel/utils/common"

type BookApply struct {
	Model
	Uid      uint   `json:"uid,omitempty"`
	Name     string `json:"name,omitempty"`
	Author   string `json:"author,omitempty"`
	Platform string `json:"platform,omitempty"`
	Reply    string `json:"reply,omitempty"`
	BookId   uint   `json:"book_id,omitempty"`
	ReplyAt  string `json:"reply_at,omitempty"`
}

func (m BookApply) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []BookApply {
	list := make([]BookApply, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m BookApply) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}

func (m Book) Insert() uint {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}
