package timetool

import (
	"time"

	"github.com/jinzhu/now"
)

// 本月开始
func BeginOfMonth() time.Time {
	return now.BeginningOfMonth()
}

// 上月开始
func BeginOfLastMonth() time.Time {
	return BeginOfMonth().AddDate(0, -1, 0)
}

// 本月结束日期
func EndDayOfMonth() time.Time {
	end := now.EndOfMonth()
	return NewDay(end.Year(), end.Month(), end.Day())
}

// 本月结束时间
func EndTimeOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Second)
}

// 上月结束日期
func EndDayOfLastMonth() time.Time {
	return BeginOfMonth().AddDate(0, 0, -1)
}

// 上月结束时间
func EndTimeOfLastMonth() time.Time {
	return BeginOfMonth().Add(-1 * time.Second)
}

// 本月日期时间段
func DurationDayOfMonth() (time.Time, time.Time) {
	return BeginOfMonth(), EndDayOfMonth()
}

// 本月时间段
func DurationTimeOfMonth() (time.Time, time.Time) {
	return BeginOfMonth(), EndTimeOfMonth()
}

// 上月日期时间段
func DurationDayOfLastMonth() (time.Time, time.Time) {
	return BeginOfLastMonth(), EndDayOfLastMonth()
}

// 上月时间段
func DurationTimeOfLastMonth() (time.Time, time.Time) {
	return BeginOfLastMonth(), EndTimeOfLastMonth()
}
