package base

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/jdongdong/go-lib/slog"
	"github.com/pquerna/ffjson/ffjson"
)

type BaseController struct {
	beego.Controller
	BaseControllerInterface
}

type BaseControllerInterface interface {
}

func (this *BaseController) ReqLog() {
	req := make(map[string]interface{})
	ffjson.Unmarshal(this.Ctx.Input.RequestBody, &req)
	vlu, _ := ffjson.Marshal(req)
	slog.Trace(fmt.Sprintf("url:%s body:%s", this.Ctx.Request.RequestURI, string(vlu)))
}

func (this *BaseController) ToIntEx(s string, defaultVlu ...int) int {
	if rs, err := this.GetInt(s); err == nil {
		return rs
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (this *BaseController) ToInt64Ex(s string, defaultVlu ...int64) int64 {
	if rs, err := this.GetInt64(s); err == nil {
		return rs
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (this *BaseController) NeedAuth(s ...string) bool {
	uri := this.Ctx.Request.RequestURI
	for _, v := range s {
		if strings.Contains(uri, v) {
			return false
		}
	}
	return true
}
