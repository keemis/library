package timer

import (
	"time"
)

const GoTimeFormat = "2006-01-02 15:04:05"

// FormatCN 常用格式
func FormatCN(t time.Time) string {
	return t.Format(GoTimeFormat)
}

// ParseCN 常用格式
func ParseCN(str string) (time.Time, error) {
	return time.ParseInLocation(GoTimeFormat, str, time.Local)
}
