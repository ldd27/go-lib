package email

import (
	"github.com/futurenda/gomail"
)

type Option struct {
	Host      string
	Port      int
	User      string
	Password  string
	EnableTLS bool
	EnableSSL bool
	Timeout   int
}

var defaultOption = Option{}

func NewEmail(opts ...func(*Option)) *gomail.Dialer {
	conf := defaultOption
	for _, o := range opts {
		o(&conf)
	}

	return gomail.NewDialer(conf.Host, conf.Port, conf.User, conf.Password)
}
