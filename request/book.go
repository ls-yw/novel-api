package request

type BookList struct {
	Pages
	Cid     uint   `form:"cid" json:"cid"`
	Keyword string `form:"keyword" json:"keyword"`
}
