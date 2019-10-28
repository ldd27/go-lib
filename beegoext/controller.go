package beegoext

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/ldd27/go-lib/slog"

	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
	BaseControllerInterface
}

type BaseControllerInterface interface {
}

func (r *Controller) ReqLog() {
	req := make(map[string]interface{})
	jsoniter.Unmarshal(r.Ctx.Input.RequestBody, &req)
	vlu, _ := jsoniter.Marshal(req)
	slog.Trace(fmt.Sprintf("url:%s body:%s", r.Ctx.Request.RequestURI, string(vlu)))
}

func (r *Controller) ToIntEx(s string, defaultVlu ...int) int {
	if rs, err := r.GetInt(s); err == nil {
		return rs
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (r *Controller) ToInt64Ex(s string, defaultVlu ...int64) int64 {
	if rs, err := r.GetInt64(s); err == nil {
		return rs
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (r *Controller) NeedAuth(s ...string) bool {
	uri := r.Ctx.Request.RequestURI
	for _, v := range s {
		if strings.Contains(uri, v) {
			return false
		}
	}
	return true
}
