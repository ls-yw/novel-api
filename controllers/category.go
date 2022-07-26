package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/servers"
	"novel/utils/errors"
	"novel/woodlsy"
)

type Category struct {
	basic.Controller
}

func (ca Category) All(c *gin.Context) {

	a := woodlsy.Configs
	fmt.Printf("%+v", a)

	d := servers.GetOne()

	b := map[string]interface{}{
		"data": a,
		"row":  d,
	}

	errors.Success.ReturnJson(c, b)
}
