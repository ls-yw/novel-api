package models

import "novel/app/utils/common"

type SmsLog struct {
	Model
	Mobile    string `json:"mobile,omitempty"`
	Content   string `json:"content,omitempty"`
	IsSuccess int8   `json:"is_success,omitempty"`
	Result    string `json:"result,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Action    string `json:"action,omitempty"`
}

func (m SmsLog) Insert() int {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}
