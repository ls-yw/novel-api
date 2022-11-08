package controllers

import (
	"github.com/gin-gonic/gin"
	request2 "github.com/woodlsy/woodGin/request"
	"novel/app/data/struct/request"
	"novel/app/servers"
	"novel/app/utils/errors"
)

type Article struct {
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

func (a Article) Info(c *gin.Context) {
	var params request.ArticleInfo
	_ = c.ShouldBindQuery(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	if params.Id == 0 {
		params.Id = servers.GetFirstArticleIdByBook(params.BookId)
	}

	data := make(errors.Data)

	if params.Id == 0 {
		errors.NoFoundArticle.ReturnJson(c)
	}

	article := servers.GetArticleInfo(params.BookId, params.Id, "id,title,sort,is_oss")
	if article.Id == 0 {
		data["errorContent"] = "该章节不存在"
	} else if article.IsOss == 0 {
		data["errorContent"] = "章节内容暂未上传"
	} else {
		data["content"] = servers.GetArticleContent(params.BookId, article.Id)
	}

	data["article"] = article
	data["prevId"] = servers.GetArticlePrev(params.BookId, article.Sort)
	data["nextId"] = servers.GetArticleNext(params.BookId, article.Sort)

	errors.Success.ReturnJson(c, data)
}
