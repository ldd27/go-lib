package util

import (
	"testing"

	"github.com/ldd27/go-lib/zaplog"

	"github.com/ldd27/go-lib/log"
)

func init() {
	log.InitLog()
}

func TestNewAliyunSms(t *testing.T) {
	client, err := NewAliyunSms(func(option *Option) {
	})
	if err != nil {
		t.Error(err)
	}

	err = client.Send("13000000000", "xx", map[string]interface{}{})
	if err != nil {
		zaplog.WithError(err).Error("send err")
		t.Errorf("%+v", err)
	}
}
