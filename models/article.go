package models

type Article struct {
	Model
	Title      string `json:"title,omitempty"`
	ChapterId  int    `json:"chapterId,omitempty"`
	BookId     int    `json:"bookId,omitempty"`
	Sort       int16  `json:"sort,omitempty"`
	Wordnumber int    `json:"wordnumber,omitempty"`
	Url        string `json:"url,omitempty"`
	IsOss      int8   `json:"is_oss,omitempty"`
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
