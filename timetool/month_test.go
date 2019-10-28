package timetool

import (
	"fmt"
	"testing"
)

// 本月日期时间段
func TestDurationDayOfMonth(t *testing.T) {
	begin, end := DurationDayOfMonth()
	fmt.Println(begin, end)
}

// 本月时间段
func TestDurationTimeOfMonth(t *testing.T) {
	begin, end := DurationTimeOfMonth()
	fmt.Println(begin, end)
}

// 上月日期时间段
func TestDurationDayOfLastMonth(t *testing.T) {
	begin, end := DurationDayOfLastMonth()
	fmt.Println(begin, end)
}

// 上月时间段
func TestDurationTimeOfLastMonth(t *testing.T) {
	begin, end := DurationTimeOfLastMonth()
	fmt.Println(begin, end)
}
