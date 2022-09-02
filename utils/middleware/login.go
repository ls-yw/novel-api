package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel/utils/common"
	"novel/utils/errors"
	"strings"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		tokenArr := strings.Split(token, ".")
		if len(tokenArr) != 3 {
			c.JSON(http.StatusOK, errors.NoLogin)
			c.Abort()
			return
		}
		_, err := common.ParseJwt(token)
		if err != nil {
			c.JSON(http.StatusOK, errors.NoLogin)
			c.Abort()
			return
		}
		c.Next()
	}
}
