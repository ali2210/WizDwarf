package date_time

import (
	"fmt"
	"time"
)

// DATE PARSE
func Date(s string) string {
	var d string
	if (s[2:3]) == "-" {
		d = s[0:2]
	} else if s[1:2] == "-" {
		d = s[0:1]
	}
	return d
}

// MONTH PARSE
func Month(s string) string {
	var m string
	fmt.Println("month-1", s[2:4], "month-2", s[3:5])
	if s[2:3] == "-" {
		m = s[3:4]
	} else {
		m = s[3:5]
	}
	return m
}

// YEAR PARSE
func Year(s string) string {
	var y string
	if len(s) == 10 {
		y = s[6:10]
	} else if len(s) == 9 {
		y = s[5:9]
	}
	return y
}

func GetToday(year int, month time.Month, date int) time.Time {
	now := time.Now()
	return time.Date(year, month, date, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
}
