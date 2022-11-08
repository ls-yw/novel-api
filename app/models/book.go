package models

type Book struct {
	Model
	Name              string `json:"name"`
	Category          int    `json:"category,omitempty"`
	Author            string `json:"author,omitempty"`
	Intro             string `json:"intro,omitempty"`
	ThumbImg          string `json:"thumb_img,omitempty"`
	Click             int    `json:"click,omitempty"`
	Monthclick        int16  `json:"monthclick,omitempty"`
	Weekclick         int16  `json:"weekclick,omitempty"`
	Dayclick          int16  `json:"dayclick,omitempty"`
	Recommend         int    `json:"recommend,omitempty"`
	Coll              int16  `json:"coll,omitempty"`
	IsFinished        int8   `json:"is_finished,omitempty"`
	Articlenum        int8   `json:"articlenum,omitempty"`
	Wordsnumber       int    `json:"wordsnumber,omitempty"`
	CollectId         int    `json:"collect_id,omitempty"`
	FromCollectBookId string `json:"from_collect_book_id,omitempty"`
	IsCollect         int8   `json:"is_collect,omitempty"`
	LastCollectAt     string `json:"last_collect_at,omitempty"`
	LastAt            string `json:"last_at,omitempty"`
	IsRecommend       int8   `json:"is_recommend,omitempty"`
	Quality           int8   `json:"quality,omitempty"`
}

func (m Book) GetOne(where map[string]interface{}, orderBy string, fields string) (info Book) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m Book) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []Book {
	list := make([]Book, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m Book) GetAll(where map[string]interface{}, orderBy string, fields string) []Book {
	list := make([]Book, 0)
	getAll(&list, where, orderBy, fields)
	return list
}

func (m Book) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}
