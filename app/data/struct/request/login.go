package request

type LoginSendSmsCode struct {
	ImgCode   string `form:"imgCode" verify:"required"`
	Mobile    string `form:"mobile" verify:"required;eq=11"`
	ImgCodeId string `form:"imgCodeId" verify:"required"`
}

type LoginRegisterForm struct {
	Mobile     string `form:"mobile" verify:"required;eq=11"`
	Password   string `form:"password" verify:"required;ge=6"`
	RePassword string `form:"rePassword" verify:"required;ge=6"`
	SmsCode    string `form:"smsCode" verify:"required;eq=6"`
	Username   string `form:"username" verify:"required;ge=6;le=12"`
}

type LoginForm struct {
	Username string `form:"username" verify:"required;ge=6"`
	Password string `form:"password" verify:"required;ge=6"`
}
