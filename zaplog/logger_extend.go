package zaplog

import (
	"github.com/goadesign/goa"
	"go.uber.org/zap"
)

type ExtendLogger struct {
	*zap.SugaredLogger
}

func (r *ExtendLogger) New(keyvals ...interface{}) goa.LogAdapter {
	if r.SugaredLogger == nil {
		return r
	}
	r.SugaredLogger.With(keyvals...)
	return r
}

func (r *ExtendLogger) Info(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Infow(msg, keyvals...)
}

func (r *ExtendLogger) Error(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Errorw(msg, keyvals...)
}

func (r *ExtendLogger) Print(v ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Info(v...)
}
