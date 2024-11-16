package util

import (
	"time"
)

const (
	// TimeFormat 标准日期时间格式
	TimeFormat = "2006-01-02 15:04:05"

	// DateFormat 标准日期格式
	DateFormat      = "2006-01-02"
	PDateHourFormat = "2006-01-02 15"
)

func DateToTimestamp(dateStr string) (int64, error) {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return 0, err
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func TimestampToDate(t int64) string {
	return time.Unix(t, 0).Format(DateFormat)
}
