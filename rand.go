package gutil

import (
	"time"
	"math/rand"
)

func RandIntn(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}
