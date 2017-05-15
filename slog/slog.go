package slog

import "github.com/cihub/seelog"

func init() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		seelog.Critical("log config not exist")
		return
	}
	logger.SetAdditionalStackDepth(1)
	seelog.ReplaceLogger(logger)
}

func Trace(v ...interface{}) {
	seelog.Trace(v)
}

func Debug(v ...interface{}) {
	seelog.Debug(v)
}

func Info(v ...interface{}) {
	seelog.Info(v)
}

func Warn(v ...interface{}) {
	seelog.Warn(v)
}

func Error(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if v[0] == nil {
		return
	}
	seelog.Error(v)
}

func Critical(v ...interface{}) {
	if len(v) == 0 {
		return
	}
	if v[0] == nil {
		return
	}
	seelog.Critical(v)
}
