package controllers

import (
	"github.com/gin-gonic/gin"
	request2 "github.com/woodlsy/woodGin/request"
	"novel/app/data/global"
	"novel/app/data/struct/request"
	servers2 "novel/app/servers"
	"novel/app/utils/errors"
)

type Book struct {
}

//
// GetBookListByWeekClick
// @Description: 首页周排行榜
// @receiver b
// @param c
//
func (b Book) GetBookListByWeekClick(c *gin.Context) {
	data := make(errors.Data)
	data["data"] = servers2.GetBookListByWeekClick()
	errors.Success.ReturnJson(c, data)
}

//
// List
// @Description: 小说列表
// @receiver b
// @param c
//
func (b Book) List(c *gin.Context) {
	var params request.BookList
	_ = c.ShouldBindQuery(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}
	if params.Page == 0 {
		params.Page = 1
	}
	if params.Size == 0 {
		params.Size = 20
	}
	fields := "id,name,thumb_img,author,intro"
	data := make(errors.Data)
	data["list"] = servers2.GetBookList(params.Cid, params.Keyword, params.Page, params.Size, fields)
	data["totalCount"] = servers2.GetBookListCount(params.Cid, params.Keyword)

	errors.Success.ReturnJson(c, data)
}

//
// Info
// @Description: 详情
// @receiver b
// @param c
//
func (b Book) Info(c *gin.Context) {
	var params request.Id
	_ = c.ShouldBindQuery(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	data := make(errors.Data)
	data["data"] = servers2.GetBookInfo(params.Id)
	data["read"] = map[string]interface{}{"id": 0}
	data["user"] = false
	if global.Uid != 0 {
		userBook := servers2.GetUserBookByBookId(global.Uid, params.Id)
		if userBook.Id != 0 {
			data["user"] = true
		}
	}

	// TODO 点击量
	errors.Success.ReturnJson(c, data)
}
