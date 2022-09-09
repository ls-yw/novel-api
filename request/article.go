package request

type ArticleList struct {
	Pages
	BookId uint `form:"book_id" verify:"required"`
}
