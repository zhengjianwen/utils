package tools

import (
	"time"
)

const base_format = "2006-01-02 15:04:05"
const date  = "2006-01-02"
const ftime  = "15:04:05"

// 获取当前时间戳
func NowTimestamp() int64 {
	return time.Now().Unix()
}

func NowDateFormat() string {
	return time.Now().Format(base_format)
}
// 当前日期 2006-01-02
func NowDateDate() string {
	return time.Now().Format(date)
}
// 当前时间 15:04:05
func NowDateTime() string {
	return time.Now().Format(ftime)
}
// 时间字符串转换时间戳
func DateTimeToTimestamp(t string) int64 {
	tobj,err := time.Parse(base_format,t)
	if err != nil{
		return 0
	}
	return tobj.Unix()
}
// 时间字符串转换日期时间戳
func DateTimeToDate(t string) int64 {
	tobj,err := time.Parse(date,t)
	if err != nil{
		return 0
	}
	return tobj.Unix()
}
// 时间戳转时间 2019-04-08 08:00:00
func TimestampToFormat(data int64) string {
	date_time := time.Unix(data, 0)
	str_time := date_time.Format(base_format)
	return str_time
}
// 时间戳转日期
func TimestampToDate(data int64) string {
	date_time := time.Unix(data, 0)
	str_time := date_time.Format(date)
	return str_time
}
// 时间戳转日期
func TimestampToTime(data int64) string {
	date_time := time.Unix(data, 0)
	str_time := date_time.Format(ftime)
	return str_time
}

func TimeToTimestamp(t time.Time) int64 {
	return t.Unix()
}

func TimeToFormat(t time.Time) string {
	return t.Format(base_format)
}

func TimeToDate(t time.Time) string {
	return t.Format(date)
}

func TimeToTime(t time.Time) string {
	return t.Format(ftime)
}

func DateTimeID() string {
	return time.Now().Format("20060102150405")
}

func DateID() string {
	return time.Now().Format("20060102")
}

func TimeID() string {
	return time.Now().Format("150405")
}

func Year() int64 {
	return int64(time.Now().Year())
}

func Month() int64 {
	return int64(time.Now().Month())
}

func Day() int64 {
	return int64(time.Now().Day())
}

func Hour() int64 {
	return int64(time.Now().Hour())
}

func Minute() int64 {
	return int64(time.Now().Minute())
}

func Second() int64 {
	return int64(time.Now().Second())
}