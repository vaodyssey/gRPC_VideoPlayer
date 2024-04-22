package utils

import (
	"time"
)

func SecondsToTimeString(second int, format string) string {
	milliseconds := second * 1000
	var t time.Time // Zero time
	t = t.Add(time.Duration(milliseconds) * time.Millisecond)
	result := t.Format(format)
	return result
}
func GetEndTime(startTime int, duration int, format string) string {
	milliseconds := (startTime + duration) * 1000
	var t time.Time // Zero time
	t = t.Add(time.Duration(milliseconds) * time.Millisecond)
	result := t.Format(format)
	return result
}
