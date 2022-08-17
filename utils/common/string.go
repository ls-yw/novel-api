package common

import "strings"

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
