package yzm

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store base64Captcha.Store = RedisStore{}

var driver base64Captcha.Driver

type Config struct {
	// 验证码宽度
	Width int
	// 验证码高度
	Height int
	// 验证码噪点 | 圆点
	Noise int
	// 验证码干扰线
	Line int
	// 验证码长度
	Length int
	// 验证码背景颜色
	BgColor *color.RGBA
	// 字体
	Fonts []string
}

func StringYzm(config Config) (string, string) {
	//配置验证码的参数
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		//BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 124},
	}
	if config.Width > 0 {
		driverString.Width = config.Width
	}
	if config.Height > 0 {
		driverString.Height = config.Height
	}
	if config.Noise > 0 {
		driverString.NoiseCount = config.Noise
	}
	if config.Line > 0 {
		driverString.ShowLineOptions = config.Line
	}
	if config.Length > 0 {
		driverString.Length = config.Length
	}
	if len(config.Fonts) > 0 {
		driverString.Fonts = config.Fonts
	}
	//if config.BgColor.A > 0 {
	//	driverString.Fonts = config.Fonts
	//}

	//ConvertFonts 按名称加载字体
	driver = driverString.ConvertFonts()

	return generateCaptcha()
}

func AudioYzm() (string, string) {
	//配置验证码的参数
	driver = &base64Captcha.DriverAudio{
		Language: "zh",
		Length:   4,
	}
	return generateCaptcha()
}

func MathYzm(config Config) (string, string) {
	driverMath := base64Captcha.DriverMath{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
	}
	if config.Width > 0 {
		driverMath.Width = config.Width
	}
	if config.Height > 0 {
		driverMath.Height = config.Height
	}
	if config.Noise > 0 {
		driverMath.NoiseCount = config.Noise
	}
	if config.Line > 0 {
		driverMath.ShowLineOptions = config.Line
	}
	if len(config.Fonts) > 0 {
		driverMath.Fonts = config.Fonts
	}
	driver = driverMath.ConvertFonts()
	return generateCaptcha()
}

func ChineseYzm(config Config) (string, string) {
	driverChinese := base64Captcha.DriverChinese{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "我们是一个超人的起点啊哦额有",
	}
	if config.Width > 0 {
		driverChinese.Width = config.Width
	}
	if config.Height > 0 {
		driverChinese.Height = config.Height
	}
	if config.Noise > 0 {
		driverChinese.NoiseCount = config.Noise
	}
	if config.Line > 0 {
		driverChinese.ShowLineOptions = config.Line
	}
	if config.Length > 0 {
		driverChinese.Length = config.Length
	}
	if len(config.Fonts) > 0 {
		driverChinese.Fonts = config.Fonts
	}
	driver = driverChinese.ConvertFonts()

	return generateCaptcha()
}

func DigitYzm(config Config) (string, string) {
	driverDigit := base64Captcha.DriverDigit{
		Height:   40,
		Width:    100,
		Length:   4,
		MaxSkew:  0,
		DotCount: 100,
	}

	if config.Width > 0 {
		driverDigit.Width = config.Width
	}
	if config.Height > 0 {
		driverDigit.Height = config.Height
	}
	if config.Length > 0 {
		driverDigit.Length = config.Length
	}
	if config.Noise > 0 {
		driverDigit.DotCount = config.Noise
	}

	driver = &driverDigit

	return generateCaptcha()
}

func generateCaptcha() (string, string) {
	//创建 Captcha
	captcha := base64Captcha.NewCaptcha(driver, store)
	//Generate 生成随机 id、base64 图像字符串
	id, b64s, _ := captcha.Generate()
	return id, b64s
}
