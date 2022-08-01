package models

type Category struct {
	Model
	Name        string `json:"name"`
	ParentId    uint   `json:"parent_id"`
	SeoName     string `json:"seo_name"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Sort        uint16 `json:"sort"`
}

func (m Category) GetOne(where map[string]interface{}, orderBy string) (info Category) {
	GetOne(&info, where, orderBy)
	return
}

func (m Category) GetList(where map[string]interface{}, orderBy string, offset int, limit int) []Category {
	list := make([]Category, 0)
	GetList(&list, where, orderBy, offset, limit)
	return list
}

func (m Category) GetAll(where map[string]interface{}, orderBy string) []Category {
	list := make([]Category, 0)
	GetAll(&list, where, orderBy)
	return list
}
