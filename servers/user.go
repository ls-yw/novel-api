package servers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"novel/models"
	"novel/utils/common"
	"novel/utils/errors"
)

func GetUserByMobile(mobile string, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"mobile": mobile}, "id desc", fields)
}

func GetUserById(id int, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"id": id}, "id desc", fields)
}

func GetUserByUsername(username string, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"username": username}, "id desc", fields)
}

func UpdateUserByLogin(uid int, ip string) {
	data := map[string]interface{}{
		"last_ip":   ip,
		"last_time": common.Now(),
		"count":     gorm.Expr("count + ?", 1),
	}
	models.User{}.Update(data, map[string]interface{}{"id": uid})
}

func GetUserBookList(uid int, page int, size int) []models.UserBookList {
	offset := (page - 1) * size
	return models.UserBook{}.GetBookList(uid, offset, size)
}

func GetUserBookListCount(uid int) int64 {
	return models.UserBook{}.GetBookListCount(uid)
}

//
// AddUserBook
// @Description: 加入书架
// @param uid
// @param bookId
// @return bool
//
func AddUserBook(uid int, bookId int) bool {
	info := GetUserBookByBookId(uid, bookId)
	if info.Id != 0 {
		return true
	}
	row := models.UserBook{
		Uid:    uid,
		BookId: bookId,
	}.Insert()
	if row > 0 {
		return true
	} else {
		return false
	}
}

func GetUserBookByBookId(uid int, bookId int) models.UserBook {
	return models.UserBook{}.GetOne(map[string]interface{}{"uid": uid, "book_id": bookId}, "id desc", "id")
}

func GetUserBookById(id int) models.UserBook {
	return models.UserBook{}.GetOne(map[string]interface{}{"id": id}, "id desc", "id,uid")
}

func DelUserBook(uid int, id int, c *gin.Context) int64 {
	info := GetUserBookById(id)
	if info.Id == 0 || info.Uid != uid {
		errors.DataNoFound.ReturnJson(c)
	}
	return models.UserBook{}.Delete(map[string]interface{}{"id": id})
}

func UpdateUserBook(uid int, bookId int, articleId int) {
	info := models.UserBook{}.GetOne(map[string]interface{}{"uid": uid, "book_id": bookId}, "id desc", "id")
	if info.Id > 0 {
		models.UserBook{}.Update(map[string]interface{}{"article_id": articleId}, map[string]interface{}{"id": info.Id})
	}
}

//
// ApplyBook
// @Description: 申请收录
// @param uid
// @param bookName
// @param bookAuthor
// @return int
//
func ApplyBook(uid int, bookName string, bookAuthor string) int {
	return models.BookApply{
		Name:   bookName,
		Author: bookAuthor,
		Uid:    uid,
	}.Insert()

}

func UpdatePassword(uid int, salt string, password string) {
	models.User{}.Update(map[string]interface{}{"password": common.Md5(salt + common.Md5(password))}, map[string]interface{}{"id": uid})
}
