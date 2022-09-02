package servers

import (
	"github.com/gin-gonic/gin"
	"novel/models"
	"novel/utils/common"
	"novel/utils/errors"
	"novel/utils/redis"
	"strconv"
)

//
// Register
// @Description: 注册
// @param username
// @param mobile
// @param password
// @return uint
//
func Register(c *gin.Context, username string, mobile string, password string) uint {
	salt := common.RandString(8, 7)

	userByMobile := GetUserByMobile(mobile, "id")
	if userByMobile.Id > 0 {
		errors.MobileRegisterExists.ReturnJson(c)
	}

	userByUsername := GetUserByUsername(username, "id")
	if userByUsername.Id > 0 {
		errors.UsernameRegisterExists.ReturnJson(c)
	}

	user := models.User{
		Username: username,
		Mobile:   mobile,
		Salt:     salt,
		Password: encryptPassword(password, salt),
		LastTime: common.InitDateTime,
	}

	return user.Insert()
}

func Login(uid uint) string {
	user := GetUserById(uid, "id,username")
	token := common.Md5(user.Password)
	redis.SetEx(redis.LoginToken+token, 86400*7, strconv.Itoa(int(user.Id)))
	return common.CreateJwt(86400*7, map[string]interface{}{"token": token})
}

//
// VerifyPassword
// @Description: 验证密码
// @param password
// @param user
// @return bool
//
func VerifyPassword(password string, user models.User) bool {
	return common.Md5(user.Salt+common.Md5(password)) == user.Password
}

func encryptPassword(password string, salt string) string {
	return common.Md5(salt + common.Md5(password))
}
