package timetool

import (
	"time"

	"github.com/jinzhu/now"
)

// 本年开始
func BeginOfYear() time.Time {
	return now.BeginningOfYear()
}

// 上一年开始
func BeginOfLastYear()time.Time  {
	return now.BeginningOfYear().AddDate(-1,0,0)
}

// 本年结束日期
func EndDayOfYear() time.Time {
	end := now.EndOfYear()
	return NewDay(end.Year(), end.Month(), end.Day())
}

// 上一年结束日期
func EndDayOfLastYear()time.Time  {
	return EndDayOfYear().AddDate(-1,0,0)
}

// 本年日期时间段
func DurationDayOfYear() (time.Time, time.Time) {
	return BeginOfYear(), EndDayOfYear()
}

// 上一年日期时间段
func DurationDayOfLastYear() (time.Time, time.Time) {
	return BeginOfLastYear(), EndDayOfLastYear()
}
