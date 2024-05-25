package cmd

import (
	"fmt"
	"time"
)

func DateAndTime() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%d-%d-%d:%d",
		currentTime.Day(),
		currentTime.Month(),
		currentTime.Year(),
		currentTime.Hour(),
		currentTime.Minute())
}