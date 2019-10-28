package util

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/ldd27/go-lib/zaplog"

	jsoniter "github.com/json-iterator/go"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type Client struct {
	Client *dysmsapi.Client
	Option Option
}

type Option struct {
	RegionID        string
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
}

var defaultOption = Option{
	RegionID: "cn-hangzhou",
}

func NewAliyunSms(opts ...func(*Option)) (*Client, error) {
	conf := defaultOption
	for _, o := range opts {
		o(&conf)
	}

	client, err := dysmsapi.NewClientWithAccessKey(conf.RegionID, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		zaplog.WithError(err).Error("init sms err")
		return nil, err
	}

	res := &Client{
		Client: client,
		Option: conf,
	}
	return res, nil
}

func (r *Client) Send(phone, tplCode string, tplParam map[string]interface{}) error {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = r.Option.SignName
	request.TemplateCode = tplCode

	strTplParam, err := jsoniter.MarshalToString(tplParam)
	if err != nil {
		zaplog.WithError(err).Error("marshal sms err")
		return err
	}
	request.TemplateParam = strTplParam

	response, err := r.Client.SendSms(request)
	if err != nil {
		zaplog.WithError(err).Error("send sms err")
		return err
	}

	zaplog.WithField(zap.String("response", fmt.Sprintf("%+v", response))).Info("sms res")
	return nil
}

func (r *Client) SendWithSign(phone, signName, tplCode string, tplParam map[string]interface{}) error {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = tplCode

	strTplParam, err := jsoniter.MarshalToString(tplParam)
	if err != nil {
		zaplog.WithError(err).Error("marshal sms err")
		return err
	}
	request.TemplateParam = strTplParam

	response, err := r.Client.SendSms(request)
	if err != nil {
		zaplog.WithError(err).Error("send sms err")
		return err
	}
	zaplog.WithField(zap.String("response", fmt.Sprintf("%+v", response))).Info("sms res")
	return nil
}
