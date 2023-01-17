package sms

import "os"

type ClientInterface interface {
	SendSms(params *Params) error
}

type Params struct {
	SignName       string            // 签名
	Phone          string            // 手机号
	TemplateCode   string            // 短信模板
	TemplateParams map[string]string // 模板参数
}

func Default() ClientInterface {
	defaultClient := os.Getenv("sms.Default")
	switch defaultClient {
	case "ali":
		return NewAliCloud()
	default:
		return NewAliCloud()
	}
}
