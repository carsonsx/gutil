package gutil

import "time"

func ToPerSecond(total int64, duration time.Duration) int64 {
	return int64(float64(total) / float64(duration) * float64(time.Second))
}
