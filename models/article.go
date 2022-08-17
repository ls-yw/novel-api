package models

type Article struct {
	Model
	Title      string `json:"title,omitempty"`
	ChapterId  uint   `json:"chapterId,omitempty"`
	BookId     uint   `json:"bookId,omitempty"`
	Sort       uint16 `json:"sort,omitempty"`
	wordnumber uint   `json:"wordnumber,omitempty"`
	url        string `json:"url,omitempty"`
	is_oss     uint8  `json:"is_Oss,omitempty"`
}

func (m Article) GetOne(where map[string]interface{}, orderBy string, fields string) (info Article) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m Article) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []Article {
	list := make([]Article, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m Article) GetAll(where map[string]interface{}, orderBy string, fields string) []Article {
	list := make([]Article, 0)
	getAll(&list, where, orderBy, fields)
	return list
}

func (m Article) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}
