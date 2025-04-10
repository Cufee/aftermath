package models

import "time"

func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func StringToTime(str string) time.Time {
	sane, err := time.Parse(time.RFC3339Nano, str)
	if err == nil {
		return sane
	}
	return time.Time{}
}
