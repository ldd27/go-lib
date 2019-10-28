package timetool

import (
	"time"

	"github.com/jinzhu/now"
)

func init() {
	now.WeekStartDay = time.Monday
}

// 本周开始
func BeginOfWeek() time.Time {
	return now.BeginningOfWeek()
}

// 上周开始
func BeginOfLastWeek() time.Time {
	return BeginOfWeek().AddDate(0, 0, -7)
}

// 本周结束日期
func EndDayOfWeek() time.Time {
	return FormatDay(now.EndOfWeek())
}

// 本周结束时间
func EndTimeOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Second)
}

// 上周结束日期
func EndDayOfLastWeek() time.Time {
	return BeginOfLastWeek().Add(6 * 24 * time.Hour)
}

// 上周结束时间
func EndTimeOfLastWeek() time.Time {
	return EndTimeOfWeek().AddDate(0, 0, -7)
}

// 本周日期时间段
func DurationDayOfWeek() (time.Time, time.Time) {
	return BeginOfWeek(), EndDayOfWeek()
}

// 本周时间段
func DurationTimeOfWeek() (time.Time, time.Time) {
	return BeginOfWeek(), EndTimeOfWeek()
}

// 上周日期时间段
func DurationDayOfLastWeek() (time.Time, time.Time) {
	return BeginOfLastWeek(), EndDayOfLastWeek()
}

// 上周时间段
func DurationTimeOfLastWeek() (time.Time, time.Time) {
	return BeginOfLastWeek(), EndTimeOfLastWeek()
}
