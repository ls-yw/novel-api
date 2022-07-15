package models

type Category struct {
	Id       uint   `json:"id" json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentId uint   `json:"parent_id,omitempty"`
}

func (m Category) GetOne() {
	Orm.GetOne(&m, map[string]interface{}{})
}

//func (c Category) GetOne() Category {
//	driver.Orm{}.GetOne()
//	return c
//}

//func (Category) TableName() string {
//	return "category"
//}
