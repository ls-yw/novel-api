package models

type Book struct {
	Model
	Name              string `json:"name"`
	Category          uint   `json:"category,omitempty"`
	Author            string `json:"author,omitempty"`
	Intro             string `json:"intro,omitempty"`
	ThumbImg          string `json:"thumb_img,omitempty"`
	Click             uint   `json:"click,omitempty"`
	Monthclick        uint16 `json:"monthclick,omitempty"`
	Weekclick         uint16 `json:"weekclick,omitempty"`
	Dayclick          uint16 `json:"dayclick,omitempty"`
	Recommend         uint   `json:"recommend,omitempty"`
	Coll              uint16 `json:"coll,omitempty"`
	IsFinished        uint8  `json:"is_finished,omitempty"`
	Articlenum        uint8  `json:"articlenum,omitempty"`
	Wordsnumber       uint   `json:"wordsnumber,omitempty"`
	CollectId         uint   `json:"collect_id,omitempty"`
	FromCollectBookId string `json:"from_collect_book_id,omitempty"`
	IsCollect         uint8  `json:"is_collect,omitempty"`
	LastCollectAt     string `json:"last_collect_at,omitempty"`
	LastAt            string `json:"last_at,omitempty"`
	IsRecommend       uint8  `json:"is_recommend,omitempty"`
	Quality           uint8  `json:"quality,omitempty"`
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
