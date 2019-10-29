package zaplog

import (
	"github.com/goadesign/goa"
	"go.uber.org/zap"
)

type LoggerExtend struct {
	*zap.SugaredLogger
}

func (r *LoggerExtend) New(keyvals ...interface{}) goa.LogAdapter {
	if r.SugaredLogger == nil {
		return r
	}
	r.SugaredLogger.With(keyvals...)
	return r
}

func (r *LoggerExtend) Info(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Infow(msg, keyvals...)
}

func (r *LoggerExtend) Error(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Errorw(msg, keyvals...)
}

func (r *LoggerExtend) Print(v ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Info(v...)
}
