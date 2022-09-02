package common

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

//
// Join
// @Description: 拼接字符串
// @param glue 分隔符
// @param args
// @return string
//
func Join(glue string, args ...string) string {
	var build strings.Builder
	for k, s := range args {
		if k != 0 {
			build.WriteString(glue)
		}
		build.WriteString(s)
	}
	return build.String()
}

//
// RandString
// @Description: 随机生成字符串
// @param length
// @param randType 1:$str1 2:$str2 3:$str1和$str2 4:$str3 5:$str1和$str3 6:$str2和$str3 7:$str1和$str2和$str3
// @return string
//
func RandString(length int, randType int) string {
	str1 := "QWERTYUIOPASDFGHJKLZXCVBNM"
	str2 := "1234567890"
	str3 := "qwertyuiopasdfghjklzxcvbnm"
	var builder strings.Builder
	if HasBitwise(1, randType) {
		builder.WriteString(str1)
	}
	if HasBitwise(2, randType) {
		builder.WriteString(str2)
	}
	if HasBitwise(4, randType) {
		builder.WriteString(str3)
	}
	str := builder.String()

	rand.Seed(time.Now().UnixNano() + rand.Int63n(9999))
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteString(string(str[rand.Intn(len(str))]))
	}
	return b.String()
}

func HasBitwise(needle int, haystack int) bool {
	return (haystack & needle) == needle
}

//
// Md5
// @Description: md5 加密
// @param str
// @return string
//
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
