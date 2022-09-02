package config

type Aliyun struct {
	Sms AliyunSms `json:"sms"`
}

type AliyunSms struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SignName        string `json:"signName"`
}
