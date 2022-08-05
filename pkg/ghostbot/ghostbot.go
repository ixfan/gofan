package ghostbot

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
)

type GhostBot struct {
	httpClient   *req.Client
	httpRequest  *req.Request
	httpResponse GhostResponse
	httpError    error
}

type GhostResponse interface {
	IsSuccess() bool
	GetContentType() string
	Bytes() []byte
	String() string
	GetStatusCode() int
	GetHeader(key string) string
}

func NewClient() *GhostBot {
	return &GhostBot{
		httpClient: req.C(),
	}
}

//Debug debug模式
func (ghostBot *GhostBot) Debug() *GhostBot {
	ghostBot.httpClient.DevMode()
	return ghostBot
}

//UserAgent 设置UA
func (ghostBot *GhostBot) UserAgent(ua string) *GhostBot {
	ghostBot.httpClient.SetUserAgent(ua)
	return ghostBot
}

//RandomUserAgent 随机UA
func (ghostBot *GhostBot) RandomUserAgent() *GhostBot {
	return ghostBot
}

//SetHeaders 设置请求头信息
func (ghostBot *GhostBot) SetHeaders(headers map[string]string) *GhostBot {
	ghostBot.httpClient.SetCommonHeaders(headers)
	return ghostBot
}

func (ghostBot *GhostBot) request() *req.Request {
	if ghostBot.httpRequest == nil {
		ghostBot.httpRequest = ghostBot.httpClient.R()
	}
	return ghostBot.httpRequest
}

//SetBodyString 设置body
func (ghostBot *GhostBot) SetBodyString(body string) *GhostBot {
	ghostBot.request().SetBodyString(body)
	return ghostBot
}

//SetQuery 设置url参数
func (ghostBot *GhostBot) SetQuery(params map[string]string) *GhostBot {
	ghostBot.request().SetQueryParams(params)
	return ghostBot
}

//SetForm 设置form data
func (ghostBot *GhostBot) SetForm(formData map[string]string) *GhostBot {
	ghostBot.request().SetFormData(formData)
	return ghostBot
}

//Get get请求
func (ghostBot *GhostBot) Get(url string) *GhostBot {
	response, err := ghostBot.request().Get(url)
	ghostBot.httpResponse = response
	ghostBot.httpError = err
	return ghostBot
}

//Post post请求
func (ghostBot *GhostBot) Post(url string) *GhostBot {
	response, err := ghostBot.request().Post(url)
	ghostBot.httpResponse = response
	ghostBot.httpError = err
	return ghostBot
}

//Response 返回数据
func (ghostBot *GhostBot) Response() (GhostResponse, error) {
	return ghostBot.httpResponse, ghostBot.httpError
}

//String 返回结果字符串
func (ghostBot *GhostBot) String() (string, error) {
	if ghostBot.httpError != nil {
		return "", ghostBot.httpError
	}
	return ghostBot.httpResponse.String(), nil
}

//Bytes 返回结果字节数组
func (ghostBot *GhostBot) Bytes() ([]byte, error) {
	if ghostBot.httpError != nil {
		return nil, ghostBot.httpError
	}
	return ghostBot.httpResponse.Bytes(), nil
}

//ToStruct 返回结果结构体
func (ghostBot *GhostBot) ToStruct(value interface{}) error {
	if ghostBot.httpError != nil {
		return ghostBot.httpError
	}
	err := json.Unmarshal(ghostBot.httpResponse.Bytes(), &value)
	return err
}

func (ghostBot *GhostBot) document() (*goquery.Document, error) {
	response, err := ghostBot.Response()
	if err != nil {
		return nil, err
	}
	document, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Bytes()))
	if err != nil {
		return nil, err
	}
	return document, nil
}

//QueryFirst 查找第一个节点
func (ghostBot *GhostBot) QueryFirst(selector string) (*goquery.Selection, error) {
	doc, err := ghostBot.document()
	if err != nil {
		return nil, err
	}
	return doc.Find(selector).First(), nil
}

//QueryFind 查找多个节点
func (ghostBot *GhostBot) QueryFind(selector string) ([]*goquery.Selection, error) {
	selections := make([]*goquery.Selection, 0)
	doc, err := ghostBot.document()
	if err != nil {
		return selections, err
	}
	doc.Find(selector).Each(func(i int, selection *goquery.Selection) {
		selections = append(selections, selection)
	})
	return selections, nil
}
