package util


import (
	"github.com/google/uuid"
	"time"
)

func GetUUID() string{
	return uuid.New().String()
}

func ParseDate(date string) time.Time {
	layout := "2006-01-02"
	res, error := time.Parse(layout, date)
	if error != nil {
		panic(error)
	}
	return res
}