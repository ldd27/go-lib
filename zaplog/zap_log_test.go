package zaplog

import (
	"testing"

	"go.uber.org/zap"
)

func TestLog(t *testing.T) {
	Debug("test debug")
	Info("test info")
	Warn("test warn")
	Error("test error")
	DPanic("test dpanic")

	Logger().With(zap.String("xxx", "ddddd")).Info("dddd")
	WithField(zap.String("xx", "dd")).Info("4444")
	GetExtendLogger().Info("xx", "ctrl", "Swagger", "files", "public/swagger/swagger.json", "route", "GET /swagger.json")
	SugarLogger().Infow("xx", "ctrl", "Swagger", "files", "public/swagger/swagger.json", "route", "GET /swagger.json")

	GetExtendLogger().Print("xxx", "ddd")
}
