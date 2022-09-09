package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/utils/common"
	"novel/utils/errors"
	"novel/utils/global"
	"novel/utils/redis"
	"strconv"
	"strings"
)

func GetLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		tokenArr := strings.Split(token, ".")

		if len(tokenArr) != 3 {
			c.Next()
			return
		}
		jwtData, err := common.ParseJwt(token)
		if err != nil {
			c.Next()
			return
		}
		redisKey := jwtData["token"]
		if redisKey == nil {
			c.Next()
			return
		}
		memberId := redis.Get(redis.LoginToken + redisKey.(string))
		if memberId == "" {
			c.Next()
			return
		}
		mid, _ := strconv.ParseUint(memberId, 10, 64)
		global.Uid = uint(mid)
		c.Next()
	}
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.Uid == 0 {
			c.JSON(http.StatusOK, errors.NoLogin)
			c.Abort()
			return
		}

		c.Next()
	}
}
