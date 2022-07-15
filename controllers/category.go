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

	servers.GetOne()

	b := map[string]interface{}{
		"data": a,
	}

	errors.Success.ReturnJson(c, b)
	fmt.Println("============")
}
