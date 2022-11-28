package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/app/data/global"
	"novel/app/utils/errors"
	"strconv"
	"time"
)

func VerifyTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.Uid = 0

		timestamp := c.Request.Header.Get("timestamp")

		timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
		timeUnix := time.Now().Unix()

		if (timeUnix - timestampInt/1000) > 10 {
			c.JSON(http.StatusOK, errors.Invalid)
			c.Abort()
			return
		}

		c.Next()
	}
}
