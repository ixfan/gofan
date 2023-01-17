package sms

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

type AliCloud struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewAliCloud() *AliCloud {
	return &AliCloud{
		AccessKeyId:     os.Getenv("sms.AccessKeyId"),
		AccessKeySecret: os.Getenv("sms.AccessKeySecret"),
	}
}

// CreateClient 使用AK&SK初始化账号Client
func (aliCloud *AliCloud) CreateClient() (*dysmsapi.Client, error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: &aliCloud.AccessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: &aliCloud.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	result := &dysmsapi.Client{}
	result, err := dysmsapi.NewClient(config)
	return result, err
}

// SendSms 发送短信验证码
func (aliCloud *AliCloud) SendSms(params *Params) error {
	client, err := aliCloud.CreateClient()
	if err != nil {
		return err
	}
	templateParamsBytes, _ := json.Marshal(params.TemplateParams)
	resp, err := client.SendSms(&dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(params.Phone),
		SignName:      tea.String(params.SignName),
		TemplateCode:  tea.String(params.TemplateCode),
		TemplateParam: tea.String(string(templateParamsBytes)),
	})
	if err != nil {
		return err
	}
	statusCode := tea.Int32Value(resp.StatusCode)
	bodyCode := tea.StringValue(resp.Body.Code)
	if statusCode == 200 && bodyCode == "OK" {
		return nil
	}
	return errors.New("发送短信失败")
}
