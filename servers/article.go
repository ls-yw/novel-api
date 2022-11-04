package servers

import (
	"fmt"
	"novel/models"
	"novel/utils/http"
	"novel/utils/redis"
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
func GetArticleList(bookId int, page int, size int, fields string) []models.Article {
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
func GetArticleListCount(bookId int) int64 {
	where := map[string]interface{}{}
	where["book_id"] = bookId
	return models.Article{}.GetCount(where)
}

//
// GetFirstArticleIdByBook
// @Description: 获取小说中第一章的ID
// @param bookId
// @return int
//
func GetFirstArticleIdByBook(bookId int) int {
	info := models.Article{}.GetOne(map[string]interface{}{"book_id": bookId}, "sort asc", "id")
	return info.Id
}

//
// GetArticleInfo
// @Description: 获取章节详情
// @param bookId
// @param id
// @param fields
// @return models.Article
//
func GetArticleInfo(bookId int, id int, fields string) models.Article {
	return models.Article{}.GetOne(map[string]interface{}{"book_id": bookId, "id": id}, "", fields)
}
func GetArticlePrev(bookId int, sort int16) int {
	info := models.Article{}.GetOne(map[string]interface{}{"book_id": bookId, "sort": []interface{}{"<", sort}}, "sort desc", "id")
	return info.Id
}

func GetArticleNext(bookId int, sort int16) int {
	info := models.Article{}.GetOne(map[string]interface{}{"book_id": bookId, "sort": []interface{}{">", sort}}, "sort asc", "id")
	return info.Id
}

func GetArticleContent(bookId int, articleId int) string {
	redisKey := fmt.Sprintf("%s_%d_%d", redis.ArticleContent, bookId, articleId)
	if !redis.Exists(redisKey) {
		content := http.OssGetObject(bookId, articleId)
		if content == "" {
			return ""
		}
		redis.SetEx(redisKey, 3600, content)
	}
	return redis.Get(redisKey)
}
