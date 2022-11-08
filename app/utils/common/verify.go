package common

import "regexp"

//
// VerifyMobile
// @Description: 验证手机号
// @param mobile
// @return bool
//
func VerifyMobile(mobile string) bool {
	res, _ := regexp.MatchString(`^1[3456789]\d{9}$`, mobile)
	return res
}
