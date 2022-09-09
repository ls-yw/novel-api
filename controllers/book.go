package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/request"
	"novel/servers"
	"novel/utils/errors"
	"novel/utils/global"
	request2 "novel/woodlsy/request"
)

type Book struct {
	basic.Controller
}

//
// GetBookListByWeekClick
// @Description: 首页周排行榜
// @receiver b
// @param c
//
func (b Book) GetBookListByWeekClick(c *gin.Context) {
	data := make(errors.Data)
	data["data"] = servers.GetBookListByWeekClick()
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
	data["list"] = servers.GetBookList(params.Cid, params.Keyword, params.Page, params.Size, fields)
	data["totalCount"] = servers.GetBookListCount(params.Cid, params.Keyword)

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
	data["data"] = servers.GetBookInfo(params.Id)
	data["read"] = map[string]interface{}{"id": 0}
	data["user"] = false
	if global.Uid != 0 {
		userBook := servers.GetUserBookByBookId(global.Uid, params.Id)
		if userBook.Id != 0 {
			data["user"] = true
		}
	}

	// TODO 点击量
	errors.Success.ReturnJson(c, data)
}
