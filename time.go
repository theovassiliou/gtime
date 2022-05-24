package gtime

import (
	"math"
	"strconv"
	"time"
)

// IsToday returns true if ref is from today, false otherwise
func IsToday(ref time.Time) bool {
	return IsDaysBefore(ref, time.Now(), 0)
}

// IsYesterday returns true if ref is from today, false otherwise
func IsYesterday(ref time.Time) bool {
	return IsDaysBefore(ref, time.Now(), 1)
}

// IsDaysBefore returns true if t1 is n days before t2.
func IsDaysBefore(t1, t2 time.Time, n int) bool {

	return t1.AddDate(0, 0, n).Format("2006-01-02") == t2.Format("2006-01-02")
}

// IsDaysAfter returns true if t1 is n days after t2.
func IsDaysAfter(t1, t2 time.Time, n int) bool {

	return t1.AddDate(0, 0, -n).Format("2006-01-02") == t2.Format("2006-01-02")
}

// HFDistanceApart returns a human friendly description on how far apart (in days) is t1 from t2
// Examples:
// today, yesterday, the day before yesterday,
func HFFDistanceApart(t1, t2 time.Time) string {

	suffix := ""
	days := DaysApart(t2, t1)

	if days < -2 {
		suffix = " ago"
	}

	result := "today"
	switch {
	case days == 0:
		break
	case days == 1:
		result = "tomorrow"
	case days == 2:
		result = "day after tomorrow"
	case days == -1:
		result = "yesterday"
	case days == -2:
		result = "day before yesterday"
	default:
		result = strconv.Itoa(absInt(days)) + " days" + suffix
	}
	return result
}

// HFDistanceToday returns a human friendly description on how far apart (in days) is t from today
func HFDistanceToday(t time.Time) string {

	return HFFDistanceApart(t, time.Now())
}

// DaysApart returns the number of calendar days of which t2 is after t1. Returns negative
// value if t2 is before t1 or 0 if it is the same day
func DaysApart(t1, t2 time.Time) int {

	t1f := t1.Format("2006-01-02")
	t1n, _ := time.Parse("2006-01-02", t1f)

	t2f := t2.Format("2006-01-02")
	t2n, _ := time.Parse("2006-01-02", t2f)
	i := int(math.Round(t2n.Sub(t1n).Hours()))
	return i / 24
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// FractionOfMonthElapsed returns the fraction of the month that has past, since the first of month.
// At the first of the month it returns 0 as no day has past.
// On the last day of a month it returns a value around 0.96 as there is still one day left in the month.
// So the value 1 is never returned. The results is
// rounded to two decimal digit. Hours within now are not taken into consideration.
func FractionOfMonthElapsed(now time.Time) float64 {
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	_, _, daysOfMonth := lastOfMonth.Date()

	fraction := float64(currentDay-1) / float64(daysOfMonth)
	return math.Round(fraction*100) / 100
}
