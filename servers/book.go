package servers

import (
	"novel/models"
	"novel/utils/common"
)

//
// GetBookListByWeekClick
// @Description: 获取周排行榜小说
// @return []models.Book
//
func GetBookListByWeekClick() []models.Book {
	fields := "id,name,thumb_img,author,intro"
	return models.Book{}.GetList(map[string]interface{}{}, "weekclick desc", 0, 10, fields)
}

//
// GetBookList
// @Description: 获取小说列表
// @param categoryId
// @param keyword
// @param page
// @param size
// @param fields
// @return []models.Book
//
func GetBookList(categoryId int, keyword string, page int, size int, fields string) []models.Book {
	offset := (page - 1) * size
	where := map[string]interface{}{}
	if categoryId != 0 {
		where["category"] = categoryId
	}
	if len(keyword) > 0 {
		where["name"] = []interface{}{"like", common.Join("", "%%", keyword, "%%")}
		where["or"] = map[string]interface{}{"author": []interface{}{"like", common.Join("", "%%", keyword, "%%")}}
	}
	return models.Book{}.GetList(where, "id desc", offset, size, fields)
}

//
// GetBookListCount
// @Description: 获取列表条数
// @param categoryId
// @param keyword
// @return int64
//
func GetBookListCount(categoryId int, keyword string) int64 {
	where := map[string]interface{}{}
	if categoryId != 0 {
		where["category"] = categoryId
	}
	if len(keyword) > 0 {
		where["name"] = []interface{}{"like", common.Join("", "%%", keyword, "%%")}
		where["or"] = map[string]interface{}{"author": []interface{}{"like", common.Join("", "%%", keyword, "%%")}}
	}
	return models.Book{}.GetCount(where)
}

//
// GetBookInfo
// @Description: 获取小说详情
// @param id
// @return models.Book
//
func GetBookInfo(id int) models.Book {
	where := map[string]interface{}{"id": id}
	return models.Book{}.GetOne(where, "id asc", "id,name,thumb_img,author,intro,is_finished,wordsnumber")
}

func GetApplyBookList(page int, size int) []models.ReturnApplyList {
	offset := (page - 1) * size
	list := models.BookApply{}.GetApplyList(offset, size)
	for i := 0; i < len(list); i++ {
		username := []rune(list[i].User__username)
		list[i].User__username = common.Join("", string(username[0:2]), "***", string(username[len(username)-2:]))
	}
	return list
}

func GetApplyBookListCount() int64 {
	where := map[string]interface{}{}
	return models.BookApply{}.GetCount(where)
}
