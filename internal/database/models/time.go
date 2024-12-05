package models

import "time"

const (
	entTrashFormat = "2006-01-02 15:04:05.999999999Z07:00"
)

func TimeToString(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

func StringToTime(str string) time.Time {
	sane, err := time.Parse(time.RFC3339Nano, str)
	if err == nil {
		return sane
	}

	trash, err := time.Parse(entTrashFormat, str)
	if err != nil {
		return time.Time{}
	}
	return trash
}
