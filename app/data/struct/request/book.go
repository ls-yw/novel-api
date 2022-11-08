package request

type BookList struct {
	Pages
	Cid     int    `form:"cid" json:"cid"`
	Keyword string `form:"keyword" json:"keyword"`
}
