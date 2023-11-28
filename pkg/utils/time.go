package utils

import "time"

const defaultTimeFormat string = "2006-01-02 15:04:05"

func TimeFromStr(str string) (time.Time, error) {
	return time.Parse(defaultTimeFormat, str)
}

func TimeToStrWithFormat(t time.Time, formatStr string) string {
	return t.Format(defaultTimeFormat)
}

func TimeToStr(t time.Time) string {
	return TimeToStrWithFormat(t, defaultTimeFormat)
}

func NowTimeToStr() string {
	return TimeToStrWithFormat(time.Now(), defaultTimeFormat)
}
