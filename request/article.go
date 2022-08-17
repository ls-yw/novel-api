package request

type BookId struct {
	Pages
	BookId uint `form:"book_id" verify:"required"`
}
