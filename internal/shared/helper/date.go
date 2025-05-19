package helper

import (
	"time"
)

const (
	layout = "2006-01-02T15:04:05"
)

func StringToDate(date string) time.Time {
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return t
}
