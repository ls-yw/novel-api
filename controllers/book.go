package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/servers"
	"novel/utils/errors"
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
