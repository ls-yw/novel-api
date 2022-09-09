package controllers

import (
	"github.com/gin-gonic/gin"
	"novel/basic"
	"novel/models"
	"novel/request"
	"novel/servers"
	"novel/utils/common"
	"novel/utils/errors"
	"novel/utils/http"
	"novel/utils/redis"
	"novel/utils/yzm"
	"novel/woodlsy"
	request2 "novel/woodlsy/request"
	"regexp"
)

type Login struct {
	basic.Controller
}

//type configJsonBody struct {
//	Id            string
//	CaptchaType   string
//	VerifyValue   string
//	DriverAudio   *base64Captcha.DriverAudio
//	DriverString  *base64Captcha.DriverString
//	DriverChinese *base64Captcha.DriverChinese
//	DriverMath    *base64Captcha.DriverMath
//	DriverDigit   *base64Captcha.DriverDigit
//}

func (l Login) Yzm(c *gin.Context) {

	id, b64s := yzm.MathYzm(yzm.Config{Width: 100, Height: 32, Noise: 0, Line: 2})

	data := make(errors.Data)
	data["data"] = b64s
	data["id"] = id

	errors.Success.ReturnJson(c, data)
}

//
// SendSmsCode
// @Description: 发送短信验证码
// @receiver l
// @param c
//
func (l Login) SendSmsCode(c *gin.Context) {
	var params request.LoginSendSmsCode
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	if !common.VerifyMobile(params.Mobile) {
		errors.MobileFailed.ReturnJson(c)
	}

	if !(yzm.RedisStore{}.Verify(params.ImgCodeId, params.ImgCode, true)) {
		errors.ImgCodeFailed.ReturnJson(c)
	}

	smsCode := common.RandString(6, 2)

	redisKey := redis.LoginSmsCode + params.Mobile
	row := redis.SetEx(redisKey, 300, smsCode)
	if !row {
		errors.SmsCodeCancelFailed.ReturnJson(c)
	}

	res := http.SmsSendByAli(params.Mobile, http.SmsRegisterTemplateCode, "{\"code\":\""+smsCode+"\"}", woodlsy.Configs.Aliyun.Sms.SignName, "register", c.ClientIP())
	if !res {
		errors.SmsSendFailed.ReturnJson(c)
	}

	errors.Success.ReturnJson(c)
}

//
// Register
// @Description: 注册
// @receiver l
// @param c
//
func (l Login) Register(c *gin.Context) {
	var params request.LoginRegisterForm
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	if !common.VerifyMobile(params.Mobile) {
		errors.MobileFailed.ReturnJson(c)
	}

	if res, _ := regexp.MatchString(`^[a-zA-Z_-]{6,12}$`, params.Username); !res {
		errors.UsernameFailed.ReturnJson(c)
	}
	if len(params.Password) < 6 {
		errors.PasswordLengthFailed.ReturnJson(c)
	}
	if params.Password != params.RePassword {
		errors.RePasswordFailed.ReturnJson(c)
	}

	// 验证短信验证码
	/*	redisKey := redis.LoginSmsCode + params.Mobile

		if !redis.Exists(redisKey) {
			errors.SmsSaveCodeFailed.ReturnJson(c)
		}
		smsCode := redis.Get(redisKey)
		if smsCode != params.SmsCode {
			errors.SmsCodeFailed.ReturnJson(c)
		}*/

	uid := servers.Register(c, params.Username, params.Mobile, params.Password)
	if uid == 0 {
		errors.RegisterFailed.ReturnJson(c)
	}
	jwtToken := servers.Login(uid, c.ClientIP())
	if jwtToken == "" {
		errors.RegisterToLoginFailed.ReturnJson(c)
	}

	data := make(errors.Data)
	data["data"] = jwtToken
	errors.Success.ReturnJson(c, data)
}

func (l Login) Login(c *gin.Context) {
	var params request.LoginForm
	_ = c.ShouldBindJSON(&params)
	if err := request2.Validator(params); err != nil {
		resp := errors.ErrorCustom
		resp.Message = err.Error()
		resp.ReturnJson(c)
	}

	var user models.User
	if common.VerifyMobile(params.Username) {
		user = servers.GetUserByMobile(params.Username, "id,username,mobile,password,salt")
	} else {
		user = servers.GetUserByUsername(params.Username, "id,username,mobile,password,salt")
	}
	if user.Id == 0 {
		errors.UserExistsFailed.ReturnJson(c)
	}
	if !servers.VerifyPassword(params.Password, user) {
		errors.PasswordFailed.ReturnJson(c)
	}
	jwtToken := servers.Login(user.Id, c.ClientIP())
	if jwtToken == "" {
		errors.LoginFailed.ReturnJson(c)
	}
	data := make(errors.Data)
	data["data"] = jwtToken
	errors.Success.ReturnJson(c, data)
}
