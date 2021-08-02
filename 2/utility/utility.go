package utility

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	TimeFormat string = "2006-01-02 15:04:05" // TimeFormat is a const for default time format
	TimeGmt    int    = 7                     // TimeGmt is a alias for GMT(+7)
)

func Now(hourOffset int, format string) string {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Format(format)
}

func GenerateId() string {
	return uuid.NewV4().String()
}
