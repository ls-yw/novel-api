package request

type UserBookAdd struct {
	BookId int `json:"book_id"`
}

type ApplyBook struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}
