package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/utils/common"
	"novel/utils/errors"
	"novel/utils/redis"
)

type Member struct {
	basic.Controller
}

func (m Member) LoginInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	jwtData, err := common.ParseJwt(token)
	if err != nil {
		errors.JWTParseFailed.ReturnJson(c)
	}
	redisKey := jwtData["token"]
	if redisKey == nil {
		errors.JWTParseFailed.ReturnJson(c)
	}
	member := redis.Get(redisKey.(string))
	if member == "" {
		errors.InvalidLogin.ReturnJson(c)
	}
	fmt.Println(member)
}
