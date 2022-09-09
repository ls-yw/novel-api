package models

import (
	"novel/utils/common"
)

type UserBook struct {
	Model
	Uid       uint `json:"uid,omitempty"`
	BookId    uint `json:"book_id,omitempty"`
	ArticleId uint `json:"article_id,omitempty"`
}

type UserBookList struct {
	Id        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ThumbImg  string `json:"thumb_img,omitempty"`
	Title     string `json:"title,omitempty"`
	ArticleId uint   `json:"article_id,omitempty"`
	NewTitle  string `json:"new_title,omitempty"`
}

func (m UserBook) GetOne(where map[string]interface{}, orderBy string, fields string) (info UserBook) {
	getOne(&info, where, orderBy, fields)
	return
}

func (m UserBook) GetList(where map[string]interface{}, orderBy string, offset int, limit int, fields string) []UserBook {
	list := make([]UserBook, 0)
	getList(&list, where, orderBy, offset, limit, fields)
	return list
}

func (m UserBook) GetAll(where map[string]interface{}, orderBy string, fields string) []UserBook {
	list := make([]UserBook, 0)
	getAll(&list, where, orderBy, fields)
	return list
}

func (m UserBook) GetCount(where map[string]interface{}) int64 {
	return getCount(m, where)
}

func (m UserBook) Insert() uint {
	if m.CreatedAt == "" {
		m.CreatedAt = common.Now()
		m.UpdatedAt = m.CreatedAt
	}
	insert(&m)
	return m.Id
}

func (m UserBook) Update(data map[string]interface{}, where map[string]interface{}) int64 {
	data["updated_at"] = common.Now()
	return update(m, data, where)
}

func (m UserBook) Delete(where map[string]interface{}) int64 {
	return deleted(m, where)
}

func (m UserBook) GetBookList(uid uint, offset int, limit int) []UserBookList {
	result := make([]UserBookList, 0)
	Orm.Model(m).
		Select("nl_user_book.id,nl_user_book.article_id,nl_book.name,nl_book.thumb_img,nl_article.title,t4.title as new_title").
		Joins("left join nl_book on nl_book.id=nl_user_book.book_id").
		Joins("left join nl_article on nl_article.id=nl_user_book.article_id").
		Joins("left join nl_article as t4 on t4.book_id=nl_user_book.book_id and t4.sort = (select max(sort) from nl_article where book_id = nl_user_book.book_id)").
		Where("nl_user_book.uid", uid).
		Order("nl_user_book.updated_at desc").
		Offset(offset).Limit(limit).Scan(&result)
	return result
}

func (m UserBook) GetBookListCount(uid uint) int64 {
	var count int64
	Orm.Model(m).
		Select("nl_user_book.id,nl_user_book.article_id,nl_book.name,nl_book.thumb_img,nl_article.title,t4.title as new_title").
		Joins("left join nl_book on nl_book.id=nl_user_book.book_id").
		Joins("left join nl_article on nl_article.id=nl_user_book.article_id").
		Joins("left join nl_article as t4 on t4.book_id=nl_user_book.book_id and t4.sort = (select max(sort) from nl_article where book_id = nl_user_book.book_id)").
		Where("nl_user_book.uid", uid).
		Order("nl_user_book.updated_at desc").Count(&count)

	return count
}
