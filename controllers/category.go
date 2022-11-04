package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/request"
	"novel/servers"
	"novel/utils/errors"
	request2 "novel/woodlsy/request"
)

type Category struct {
	basic.Controller
}

//
// All
// @Description: 所有分类
// @receiver ca
// @param c
//
func (ca Category) All(c *gin.Context) {
	category := servers.GetCategoryAll()

	data := make(errors.Data)
	data["data"] = category
	errors.Success.ReturnJson(c, data)
}

func (ca Category) Info(c *gin.Context) {
	var params request.Id
	_ = c.ShouldBindQuery(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	category := servers.GetCategoryInfo(params.Id)

	data := make(errors.Data)
	data["data"] = category
	errors.Success.ReturnJson(c, data)
}
