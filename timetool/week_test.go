package timetool

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/now"
)

func init() {
	now.WeekStartDay = time.Monday
}

func TestDurationDayOfWeek(t *testing.T) {
	begin, end := DurationDayOfWeek()
	fmt.Println(begin, end)
}

func TestDurationTimeOfWeek(t *testing.T) {
	begin, end := DurationTimeOfWeek()
	fmt.Println(begin, end)
}

func TestDurationDayOfLastWeek(t *testing.T) {
	begin, end := DurationDayOfLastWeek()
	fmt.Println(begin, end)
}

func TestDurationTimeOfLastWeek(t *testing.T) {
	begin, end := DurationTimeOfLastWeek()
	fmt.Println(begin, end)
}
