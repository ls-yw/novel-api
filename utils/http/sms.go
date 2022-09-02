package http

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"novel/models"
	"novel/woodlsy"
	"novel/woodlsy/log"
)

const SmsRegisterTemplateCode = "SMS_232918959" // 注册验证码

//
// SmsSendByAli
// @Description: 发送阿里云短信
// @param merchantId
// @param mobile
// @param templateCode
// @param content
// @param signName
// @param action
// @param obj
// @param ip
// @return bool
//
func SmsSendByAli(mobile string, templateCode string, content string, signName string, action string, ip string) bool {
	accessKeyId := woodlsy.Configs.Aliyun.Sms.AccessKeyId
	accessKeySecret := woodlsy.Configs.Aliyun.Sms.AccessKeySecret

	data := models.SmsLog{
		Mobile:    mobile,
		Content:   content,
		IsSuccess: 0,
		Ip:        ip,
		Action:    action,
	}
	response := smsApi(accessKeyId, accessKeySecret, mobile, templateCode, content, signName)
	if response.IsSuccess() {
		data.IsSuccess = 1
	} else {
	}
	data.Result = response.GetHttpContentString()
	id := data.Insert()
	if id == 0 {
		log.Logger.Error("短信发送记录保存失败", data)
		return false
	}
	return true
}

//
// smsApi
// @Description: 通过接口发送短信
// @param accessKeyId
// @param accessKeySecret
// @param mobile
// @param templateCode
// @param content
// @param signName
// @return *dysmsapi.SendSmsResponse
//
func smsApi(accessKeyId string, accessKeySecret string, mobile string, templateCode string, content string, signName string) *dysmsapi.SendSmsResponse {
	client, _ := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile       //接收短信的手机号码
	request.SignName = signName         //短信签名名称
	request.TemplateCode = templateCode //短信模板ID
	request.TemplateParam = content     //短信模板参数
	response, _ := client.SendSms(request)
	return response
}
