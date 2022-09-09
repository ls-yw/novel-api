package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/request"
	"novel/servers"
	"novel/utils/errors"
	request2 "novel/woodlsy/request"
)

type Article struct {
	basic.Controller
}

//
// List
// @Description: 获取章节列表
// @receiver a
// @param c
//
func (a Article) List(c *gin.Context) {
	var params request.ArticleList
	_ = c.ShouldBindQuery(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	data := make(errors.Data)
	data["list"] = servers.GetArticleList(params.BookId, params.Page, params.Size, "id,title")
	data["totalCount"] = servers.GetArticleListCount(params.BookId)

	errors.Success.ReturnJson(c, data)
}
