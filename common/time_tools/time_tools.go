package time_tools

import (
	"math"
	"net/http"
	"time"
)

// RFC3339ToCSTLayout convert rfc3339 value to china standard time layout
// 2020-11-08T08:18:46+08:00 => 2020-11-08 08:18:46
func RFC3339ToCSTLayout(value string) (string, error) {
	ts, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}

	return ts.In(cst).Format(CSTLayout), nil
}

// CSTLayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func CSTLayoutString() string {
	ts := time.Now()

	return ts.In(cst).Format(CSTLayout)
}

// ParseCSTInLocation 格式化时间
func ParseCSTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(CSTLayout, date, cst)
}

// CSTLayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func CSTLayoutStringToUnix(cstLayoutString string) (int64, error) {
	stamp, err := time.ParseInLocation(CSTLayout, cstLayoutString, cst)
	if err != nil {
		return 0, err
	}

	return stamp.Unix(), nil
}

// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(cst).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, cst)
}

// SubInLocation 计算时间差
func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(cst).Sub(ts).Seconds())
}

// MinuteToNano 分钟转纳秒
func MinuteToNano(minute int64) (nano int64) {
	nano = int64(time.Minute * time.Duration(minute))

	return
}

// NanoToMilli 纳秒转毫秒
func NanoToMilli(nano int64) (milli int64) {
	milli = int64(time.Duration(nano) / time.Millisecond)

	return
}
