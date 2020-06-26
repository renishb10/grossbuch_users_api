package date_utils

import "time"

const (
	apiDateFormat = "2006-01-12T15:04:05.000Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}
