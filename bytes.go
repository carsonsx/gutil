package gutil

import (
	"fmt"
	"math"
)

func BytesIndexOf(b, sub []byte) int {
	for i := 0; i < len(b)-len(sub)+1; i++ {
		found := true
		for j := 0; j < len(sub); j++ {
			if b[i+j] != sub[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func CombineBytes(one []byte, others ...[]byte) []byte {
	l := len(one)
	for _, other := range others {
		if other == nil {
			continue
		}
		l += len(other)
	}
	newBytes := make([]byte, l)
	copy(newBytes, one)
	l = len(one)
	for _, other := range others {
		if other == nil {
			continue
		}
		copy(newBytes[l:], other)
		l += len(other)
	}
	return newBytes
}

func HumanReadableByteCount(bytes int64, si bool) string {
	var unit int64 = 1024
	if si {
		unit = 1000
	}
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	exp := (int64)(math.Log(float64(bytes)) / math.Log(float64(unit)))
	pre := "KMGTPE"
	if si {
		pre = "kMGTPE"
	}
	pre = pre[exp-1 : exp]
	if !si {
		pre += "i"
	}
	return fmt.Sprintf("%.1f %sB", float64(bytes)/math.Pow(float64(unit), float64(exp)), pre)
}
