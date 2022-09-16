package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/request"
	"novel/servers"
	"novel/utils/errors"
	"novel/utils/global"
	request2 "novel/woodlsy/request"
	"time"
)

type Member struct {
	basic.Controller
}

func (m Member) LoginInfo(c *gin.Context) {
	member := servers.GetUserById(global.Uid, "id,username,last_ip,last_time")
	lastTime, _ := time.Parse("2006-01-02T15:04:05+08:00", member.LastTime)
	member.LastTime = lastTime.Format("2006-01-02 15:04:05")

	data := make(errors.Data)
	data["data"] = member
	errors.Success.ReturnJson(c, data)
}

func (m Member) Book(c *gin.Context) {
	var params request.Pages
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

	data := make(errors.Data)
	data["list"] = servers.GetUserBookList(global.Uid, params.Page, params.Size)
	data["totalCount"] = servers.GetUserBookListCount(global.Uid)

	errors.Success.ReturnJson(c, data)
}

func (m Member) DelBook(c *gin.Context) {
	var params request.Id
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}
	if params.Id == 0 {
		errors.ParamsFailed.ReturnJson(c)
	}
	errors.DeleteFailed.ReturnJson(c)
	row := servers.DelUserBook(global.Uid, params.Id, c)
	if row == 0 {
		errors.DeleteFailed.ReturnJson(c)
	}
	errors.Success.ReturnJson(c)
}

//
// AddBook
// @Description: 加入书架
// @receiver m
// @param c
//
func (m Member) AddBook(c *gin.Context) {
	var params request.UserBookAdd
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}
	if params.BookId == 0 {
		errors.ParamsFailed.ReturnJson(c)
	}

	row := servers.AddUserBook(global.Uid, params.BookId)
	if !row {
		errors.AddUserBookFailed.ReturnJson(c)
	}
	errors.Success.ReturnJson(c)
}

//
// Apply
// @Description:申请收录
// @receiver m
// @param c
//
func (m Member) Apply(c *gin.Context) {
	var params request.ApplyBook
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}
	if params.Name == "" {
		errors.ParamsFailed.ReturnJson(c)
	}
	if params.Author == "" {
		errors.ParamsFailed.ReturnJson(c)
	}
	row := servers.ApplyBook(global.Uid, params.Name, params.Author)
	if row == 0 {
		errors.SaveFailed.ReturnJson(c)
	}
	errors.Success.ReturnJson(c)
}

func (m Member) ApplyList(c *gin.Context) {
	var params request.Pages
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

	//fields := "id,name,book_id,author,reply"
	data := make(errors.Data)
	data["list"] = servers.GetApplyBookList(params.Page, params.Size)
	data["totalCount"] = servers.GetApplyBookListCount()

	errors.Success.ReturnJson(c, data)
}

func (m Member) Read(c *gin.Context) {
	var params request.ArticleInfo
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	if params.Id == 0 || params.BookId == 0 {
		errors.ParamsFailed.ReturnJson(c)
	}

	servers.UpdateUserBook(global.Uid, params.BookId, params.Id)
	errors.Success.ReturnJson(c)
}
