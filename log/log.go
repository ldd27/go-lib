package log

import (
	"github.com/ddliao/go-lib/log/gogs_log"
	"fmt"
	"os"
	"strings"
)

func InitLog(path string) {
	if path == "" {
		fmt.Println("配置错误：未配置日志路径")
		return
	}

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("配置错误：日志文件夹创建失败", err)
		return
	}

	gogs_log.Register("trace", gogs_log.NewFileWriter)
	gogs_log.Register("debug", gogs_log.NewFileWriter)
	gogs_log.Register("info", gogs_log.NewFileWriter)
	gogs_log.Register("warn", gogs_log.NewFileWriter)
	gogs_log.Register("error", gogs_log.NewFileWriter)
	gogs_log.Register("fatal", gogs_log.NewFileWriter)

	gogs_log.NewLogger(10000, "trace", getConfig(gogs_log.TRACE, false, path), gogs_log.TRACE)
	gogs_log.NewLogger(10000, "debug", getConfig(gogs_log.DEBUG, false, path), gogs_log.DEBUG)
	gogs_log.NewLogger(10000, "info", getConfig(gogs_log.INFO, true, path), gogs_log.INFO)
	gogs_log.NewLogger(10000, "warn", getConfig(gogs_log.WARN, true, path), gogs_log.WARN)
	gogs_log.NewLogger(10000, "error", getConfig(gogs_log.ERROR, false, path), gogs_log.ERROR)
	gogs_log.NewLogger(10000, "fatal", getConfig(gogs_log.FATAL, false, path), gogs_log.FATAL)

	gogs_log.NewLogger(10000, "console", fmt.Sprintf(`{"level": %d}`, gogs_log.TRACE), gogs_log.TRACE)
}

func getConfig(level int, daily bool, logPath string) string {
	if logPath == "" {
		panic("未配置日志路径")
		return ""
	}
	logPath = strings.TrimRight(logPath, "\\") + "\\\\"

	target := ""
	switch level {
	case gogs_log.TRACE:
		logPath = logPath + "trace\\\\"
		target = "trace.log"
	case gogs_log.DEBUG:
		logPath = logPath + "debug\\\\"
		target = "debug.log"
	case gogs_log.INFO:
		logPath = logPath + "info\\\\"
		target = "info.log"
	case gogs_log.WARN:
		logPath = logPath + "warn\\\\"
		target = "warn.log"
	case gogs_log.ERROR:
		logPath = logPath + "error\\\\"
		target = "error.log"
	case gogs_log.CRITICAL:
		logPath = logPath + "critical\\\\"
		target = "critical.log"
	case gogs_log.FATAL:
		logPath = logPath + "fatal\\\\"
		target = "fatal.log"
	}
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		fmt.Errorf("日志创建失败", err)
		panic(err)
		return ""
	}

	return fmt.Sprintf(`{"level":%d,"filename":"%s","rotate":%v,"maxlines":%d,"maxsize":%d,"daily":%v,"maxdays":%d}`, level, logPath+target, true, 1000000, 1<<23, daily, 7)
}

func Trace(v ...interface{}) {
	msg := fmt.Sprint(v)
	gogs_log.Trace(4, msg)
}

func Debug(v ...interface{}) {
	msg := fmt.Sprint(v)
	gogs_log.Debug(4, msg)
}

func Warn(v ...interface{}) {
	msg := fmt.Sprint(v)
	gogs_log.Warn(4, msg)
}

func Info(v ...interface{}) {
	msg := fmt.Sprint(v)
	gogs_log.Info(4, msg)
}

func Error(v ...interface{}) {
	if len(v) == 0 || v[0] == nil {
		return
	}
	//	errMsg := ""
	//	if vlu, isOk := v[0].(error); isOk {
	//		errMsg = vlu.Error()
	//	} else if vlu, isOk := v[0].(string); isOk {
	//		errMsg = vlu
	//	}
	msg := fmt.Sprint(v)
	gogs_log.Error(4, msg)
}

func Fatal(v ...interface{}) {
	if len(v) == 0 || v[0] == nil {
		return
	}
	msg := fmt.Sprint(v)
	gogs_log.Fatal(3, msg)
}
