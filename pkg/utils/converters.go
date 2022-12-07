package utils

import "time"

func IntToSeconds(number int) time.Duration {
	return time.Duration(number) * time.Second
}
