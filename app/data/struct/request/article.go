package request

type ArticleList struct {
	Pages
	BookId int `form:"book_id" verify:"required"`
}

type ArticleInfo struct {
	BookId int `form:"book_id" json:"book_id" verify:"required"`
	Id     int `form:"id"`
}
