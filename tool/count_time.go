package tool

import (
	"time"
)

func CountTime(begin time.Time) time.Duration {
	return time.Since(begin)
}
