package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/woodlsy/woodGin/log"
	request2 "github.com/woodlsy/woodGin/request"
	"novel/app/data/struct/request"
	"novel/app/servers"
	"novel/app/utils/errors"
)

type Category struct {
}

//
// All
// @Description: 所有分类
// @receiver ca
// @param c
//
func (ca Category) All(c *gin.Context) {
	category := servers.GetCategoryAll()

	log.Logger.Error("我是error")
	log.Logger.Info("我是ingo")
	log.Logger.Warn("warn")
	log.Logger.Debug("debug")
	errors.Success.ReturnJson(c)

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
