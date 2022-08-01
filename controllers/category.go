package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/utils/errors"
)

type Category struct {
	basic.Controller
}

func (ca Category) All(c *gin.Context) {

	errors.Success.ReturnJson(c)
}
