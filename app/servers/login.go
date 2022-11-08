package servers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woodlsy/woodGin/config"
	"io/ioutil"
	"net/http"
	"novel/app/models"
	"novel/app/utils/common"
	"novel/app/utils/errors"
	"novel/app/utils/redis"
	"strconv"
)

//
// Register
// @Description: 注册
// @param username
// @param mobile
// @param password
// @return int
//
func Register(c *gin.Context, username string, mobile string, password string) int {
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

func Login(uid int, ip string) string {
	user := GetUserById(uid, "id,username")
	token := common.Md5(user.Password)
	redis.SetEx(redis.LoginToken+token, 86400*7, strconv.Itoa(int(user.Id)))
	jwt := common.CreateJwt(86400*7, map[string]interface{}{"token": token})
	if jwt != "" {
		UpdateUserByLogin(uid, ip)
	}
	return jwt
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

type TempResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func TemSanLogin(username string, password string) bool {

	fmt.Println("进入api登录")

	req := `{"code":"!&87WIs!#dxoiM@^", "username": "` + username + `","password":"` + password + `"}`
	req_new := bytes.NewBuffer([]byte(req))

	url := config.Configs.Api.Login + "/member/sanLogin"
	response, err := http.Post(url, "application/json; charset=utf-8", req_new)
	if err != nil || response.StatusCode != http.StatusOK {
		return false
	}
	res, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%+v ======\n", string(res))
	var result TempResponse
	json.Unmarshal(res, &result)
	if result.Code == 0 {
		return true
	} else {
		return false
	}
}
