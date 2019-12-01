package main

import (
	"log"
	"strings"
	"time"
)

// Config is needed for configuration from json
type Config struct {
	TimeFormat string
	InFile     string
	FromTime   string
	ToTime     string
	BufferSize int
	DiffFormat bool
}

// GetTime parse time from config
func (c *Config) GetTime(timeStamp string, isStartDate bool) time.Time {
	if strings.TrimSpace(timeStamp) == "" && isStartDate {
		return time.Now().AddDate(0, 0, -1)
	}
	if strings.TrimSpace(timeStamp) == "" && !isStartDate {
		return time.Now()
	}

	tsLen := len(strings.TrimSpace(timeStamp))

	format := "2006-02-01"
	if tsLen == 10 {
		// use default format
	}

	if tsLen == 19 {
		format = "2006-02-01 15:04:05"
	}

	if tsLen == 25 {
		format = "2006-02-01T15:04:05-07:00"
	}

	retTime, err := time.Parse(format, timeStamp)
	if err != nil {
		log.Fatalf("failed to parse time %s", timeStamp)
	}

	return retTime
}
