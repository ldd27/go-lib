package zaplog

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger

	publicLogger *zap.Logger
	publicSugar  *zap.SugaredLogger

	goaLogger *GoaLogger
)

type FileOption struct {
	FileName       string
	RollingPattern string
	MaxSize        int  // 每个日志文件保存的最大尺寸 单位：M
	MaxBackups     int  // 日志文件最多保存多少个备份
	MaxAge         int  // 文件最多保存多少天
	Compress       bool // 是否压缩
}

type Option struct {
	Level      zapcore.Level
	InitField  []zap.Field
	EncodeType string // text || json
	FileOption FileOption
}

var defaultOption = Option{
	EncodeType: "text",
	Level:      zap.DebugLevel,
	FileOption: FileOption{
		RollingPattern: "0 0 0 * * *",
		MaxSize:        100,
		MaxBackups:     100,
		MaxAge:         100,
	},
}

func RegisterRolling(opt Option, fileLogger *lumberjack.Logger) {
	c := cron.New()
	err := c.AddFunc(opt.FileOption.RollingPattern, func() {
		if fileLogger == nil {
			return
		}
		fileLogger.Rotate()
	})
	if err != nil {
		panic(fmt.Sprintf("cron pattern err: %s", err))
	}
	c.Start()
}

func init() {
	InitLog()
}

func InitLog(opts ...func(*Option)) error {
	opt := defaultOption
	for _, o := range opts {
		o(&opt)
	}

	var fileLogger *lumberjack.Logger
	if opt.FileOption.FileName != "" {
		fileLogger = &lumberjack.Logger{
			Filename:   opt.FileOption.FileName,   // 日志文件路径
			MaxSize:    opt.FileOption.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: opt.FileOption.MaxBackups, // 日志文件最多保存多少个备份
			MaxAge:     opt.FileOption.MaxAge,     // 文件最多保存多少天
			Compress:   opt.FileOption.Compress,   // 是否压缩
		}
		RegisterRolling(opt, fileLogger)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(opt.Level)

	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr))
	if fileLogger != nil {
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(fileLogger))
	}

	var core zapcore.Core
	if opt.EncodeType == "json" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
			writer,                                // 打印到控制台和文件
			atomicLevel,                           // 日志级别
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig), // 编码器配置
			writer,                                   // 打印到控制台和文件
			atomicLevel,                              // 日志级别
		)
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCallerSkip(1)
	// 设置初始化字段
	filed := zap.Fields(opt.InitField...)
	// 构造日志
	logger = zap.New(core, zap.AddCaller(), caller, filed)
	sugar = logger.Sugar()

	goaLogger = &GoaLogger{sugar}

	publicLogger = zap.New(core, zap.AddCaller(), filed)
	publicSugar = logger.Sugar()

	return nil
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("01-02 15:04:05.0000"))
}

func Logger() *zap.Logger {
	return publicLogger
}

func SugarLogger() *zap.SugaredLogger {
	return publicSugar
}

func GoaLog() *GoaLogger {
	return goaLogger
}

func Debug(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Debugw(msg, keysAndValues...)
}

func Info(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Infow(msg, keysAndValues...)
}

func Warn(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Warnw(msg, keysAndValues...)
}

func Error(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Errorw(msg, keysAndValues...)
}

func DPanic(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.DPanic(msg, fields...)
}

func DPanicf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.DPanicw(msg, keysAndValues...)
}

func Panic(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Panic(msg, fields...)
}

func Panicf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Panicw(msg, keysAndValues...)
}

func Fatal(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Fatal(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	if sugar == nil {
		return
	}
	sugar.Fatalw(msg, keysAndValues...)
}

func WithError(err error) *zap.Logger {
	return publicLogger.With(zap.Error(err))
}

func WithField(field ...zap.Field) *zap.Logger {
	return publicLogger.With(field...)
}
