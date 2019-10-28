package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

// ContextHook for log the call context
type goruntineHook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// NewContextHook use to make an hook
// 根据上面的推断, 我们递归深度可以设置到5即可.
func NewGoroutineHook(levels ...logrus.Level) logrus.Hook {
	hook := goruntineHook{
		Field:  "goroutine",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

// Levels implement levels
func (hook goruntineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire implement fire
func (hook goruntineHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] =runtime.NumGoroutine()
	return nil
}
