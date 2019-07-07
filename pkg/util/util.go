package util

import (
	"time"
)

//oneDaysagoTimestamp is return start,end timestamp one day ago
func OneDaysAgoTimestamp(t time.Time) (start, end int64) {
	c := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return c.AddDate(0, 0, -1).Unix(), (c.Unix() - 1)
}
