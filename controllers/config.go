package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/servers"
	"novel/utils/errors"
)

type Config struct {
	basic.Controller
}

//
// Index
// @Description: 获取章节列表
// @receiver a
// @param c
//
func (a Config) Index(c *gin.Context) {
	//var params request.ArticleList
	//_ = c.ShouldBindQuery(&params)
	//if err := request2.Validator(params); err != nil {
	//	resp := errors.ErrorCustom
	//	resp.Message = err.Error()
	//	resp.ReturnJson(c)
	//}

	data := make(errors.Data)
	data["data"] = servers.GetConfigs()

	errors.Success.ReturnJson(c, data)
}
