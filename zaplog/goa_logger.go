package zaplog

import (
	"github.com/goadesign/goa"
	"go.uber.org/zap"
)

type GoaLogger struct {
	*zap.SugaredLogger
}

func (r *GoaLogger) New(keyvals ...interface{}) goa.LogAdapter {
	if r.SugaredLogger == nil {
		return r
	}
	r.SugaredLogger.With(keyvals...)
	return r
}

func (r *GoaLogger) Info(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Infow(msg, keyvals...)
}

func (r *GoaLogger) Error(msg string, keyvals ...interface{}) {
	if r.SugaredLogger == nil {
		return
	}
	r.SugaredLogger.Errorw(msg, keyvals...)
}
