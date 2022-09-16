package models

type Category struct {
	Model
	Name        string `json:"name"`
	ParentId    int    `json:"parent_id,omitempty"`
	SeoName     string `json:"seo_name,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	Description string `json:"description,omitempty"`
	Sort        int16  `json:"sort,omitempty"`
}

func (m Category) GetOne(where map[string]interface{}, orderBy string, fields string) (info Category) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m Category) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []Category {
	list := make([]Category, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m Category) GetAll(where map[string]interface{}, orderBy string, fields string) []Category {
	list := make([]Category, 0)
	getAll(&list, where, orderBy, fields)
	return list
}
