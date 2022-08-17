package servers

import (
	"novel/models"
)

//
// GetArticleList
// @Description: 获取列表
// @param bookId
// @param page
// @param size
// @param fields
// @return []models.Article
//
func GetArticleList(bookId uint, page int, size int, fields string) []models.Article {
	offset := (page - 1) * size
	where := map[string]interface{}{}
	where["book_id"] = bookId
	return models.Article{}.GetList(where, "id asc", offset, size, fields)
}

//
// GetArticleListCount
// @Description: 获取列表条数
// @param bookId
// @return int64
//
func GetArticleListCount(bookId uint) int64 {
	where := map[string]interface{}{}
	where["book_id"] = bookId
	return models.Article{}.GetCount(where)
}
