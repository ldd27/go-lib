package log

import (
	"fmt"
	"log/syslog"

	"github.com/sirupsen/logrus"
	. "github.com/sirupsen/logrus/hooks/syslog"
)

type Option struct {
	EnableSysLogHook bool
	Network          string
	Addr             string
	LogLevel         int
	Tag              string
}

var defaultLogConf = Option{
	EnableSysLogHook: false,
	Network:          "udp",
	Addr:             "",
	LogLevel:         int(logrus.DebugLevel),
	Tag:              "",
}

func InitLog(opts ...func(conf *Option)) {
	conf := defaultLogConf
	for _, o := range opts {
		o(&conf)
	}

	logrus.AddHook(NewContextHook())
	logrus.AddHook(NewGoroutineHook())

	logrus.SetLevel(logrus.Level(conf.LogLevel))
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true, // Seems like automatic color detection doesn't work on windows terminals
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		QuoteEmptyFields: true,
	})

	if conf.EnableSysLogHook {
		hook, err := NewSyslogHook(conf.Network, conf.Addr, syslog.Priority(conf.LogLevel), conf.Tag)
		if err != nil {
			fmt.Println("start log udp hook fail", err)
		} else {
			logrus.AddHook(hook)
			fmt.Println("start log udp hook success")
		}
	}
}

func NewLog(opts ...func(conf *Option)) *logrus.Logger {
	conf := defaultLogConf
	for _, o := range opts {
		o(&conf)
	}

	log := logrus.New()

	log.Hooks.Add(NewContextHook())
	log.Hooks.Add(NewGoroutineHook())

	log.SetLevel(logrus.Level(conf.LogLevel))
	log.Formatter = &logrus.TextFormatter{
		ForceColors:      true, // Seems like automatic color detection doesn't work on windows terminals
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		QuoteEmptyFields: true,
	}

	if conf.EnableSysLogHook {
		hook, err := NewSyslogHook(conf.Network, conf.Addr, syslog.Priority(conf.LogLevel), conf.Tag)
		if err != nil {
			fmt.Println("start log udp hook fail", err)
		} else {
			log.Hooks.Add(hook)
			fmt.Println("start log udp hook success")
		}
	}

	return log
}
