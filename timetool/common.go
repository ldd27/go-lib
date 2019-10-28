package timetool

import "time"

func FormatDay(t time.Time) time.Time {
	return NewDay(t.Year(), t.Month(), t.Day())
}

func NewDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func DaysOfMonth(year int, month time.Month) int {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			return 30

		} else {
			return 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			return 29
		} else {
			return 28
		}
	}
}
