package oss

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//https://help.aliyun.com/document_detail/59670.html?spm=a2c4g.11186623.2.13.6289342bfPtcIT#section-ygd-qxw-kfb

var (
	Client *oss.Client
)

type Option struct {
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string
}

var defaultOption = Option{}

func NewAliyunOss(opts ...func(*Option)) error {
	conf := defaultOption
	for _, o := range opts {
		o(&conf)
	}

	client, err := oss.New(
		conf.EndPoint,
		conf.AccessKeyID,
		conf.AccessKeySecret)
	return err
}

func UploadOssFile(objectName string, reader io.Reader, contentType string) (path string, err error) {
	bucket, err := Client.Bucket(conf.Conf.Aliyuncs.BucketName)
	if err != nil {
		return
	}
	// 签名直传。
	signedURL, err := bucket.SignURL(objectName, oss.HTTPPut, 3600)
	if err != nil {
		return
	}

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType(contentType),
	}

	signedURL, err = bucket.SignURL(objectName, oss.HTTPPut, 60, options...)
	if err != nil {
		return
	}

	rsp, err := bucket.DoPutObjectWithURL(signedURL, reader, options)
	defer rsp.Body.Close()
	if err == nil {
		path = fmt.Sprintf("%s/%s", conf.Conf.Aliyuncs.BaseURL, objectName)
	}
	return
}
