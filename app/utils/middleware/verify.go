package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woodlsy/woodGin/config"
	"github.com/woodlsy/woodGin/helper"
	"github.com/woodlsy/woodGin/log"
	"io/ioutil"
	"net/http"
	"novel/app/data/global"
	"novel/app/utils/common"
	"novel/app/utils/errors"
	"reflect"
	"sort"
	"strconv"
	"time"
)

func VerifyTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		global.Platform = c.Request.Header.Get("Platform")

		timestamp := c.Request.Header.Get("timestamp")

		timestampInt, _ := strconv.ParseInt(timestamp, 10, 64)
		timeUnix := time.Now().Unix()

		if (timeUnix - timestampInt/1000) > 10 {
			log.Logger.Warn("timestamp有问题：", c.ClientIP())
			c.JSON(http.StatusOK, errors.Invalid)
			c.Abort()
			return
		}

		c.Next()
	}
}

func VerifySign() gin.HandlerFunc {
	return func(c *gin.Context) {
		timestamp := c.Request.Header.Get("timestamp")
		sign := c.Request.Header.Get("sign")

		getData := c.Request.URL.Query()

		var jsonData map[string]interface{}
		jsonBody, _ := c.GetRawData()
		_ = json.Unmarshal(jsonBody, &jsonData)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonBody))
		data := map[string]interface{}{
			"timestamp": timestamp,
		}

		for key, value := range jsonData {
			valueType := reflect.TypeOf(value).Kind()
			if valueType == reflect.Slice || valueType == reflect.Map || valueType == reflect.Struct {
				continue
			}
			data[key] = value
		}
		for key, value := range getData {
			if len(key) > 2 && string(key[len(key)-2:]) == "[]" {
				continue
			}
			data[key] = value[0]
		}

		dataKeysArray := common.GetStringKeysArrayByMap(data)
		sort.Strings(dataKeysArray)

		var dataStr []string
		for _, key := range dataKeysArray {
			dataStr = append(dataStr, fmt.Sprintf("%s=%v", key, data[key]))
		}

		str := helper.Join("&", dataStr...)
		enStr := common.Md5(helper.Join("", common.Md5(str), config.Configs.App.Custom["encrypt-salt"].(string)))

		if sign != enStr {
			log.Logger.Warn("sign有问题：", c.ClientIP())
			c.JSON(http.StatusOK, errors.Invalid)
			c.Abort()
			return
		}
		c.Next()
	}
}
