package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/servers"
	"novel/utils/errors"
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
