package gutil

import (
	"testing"
	"fmt"
)

func TestHumanReadableByteCount(t *testing.T) {
	fmt.Println(HumanReadableByteCount(0, false))
	fmt.Println(HumanReadableByteCount(27, false))
	fmt.Println(HumanReadableByteCount(999, false))
	fmt.Println(HumanReadableByteCount(1000, false))
	fmt.Println(HumanReadableByteCount(1023, false))
	fmt.Println(HumanReadableByteCount(1024, false))
	fmt.Println(HumanReadableByteCount(1728, false))
	fmt.Println(HumanReadableByteCount(110592, false))
	fmt.Println(HumanReadableByteCount(7077888, false))
	fmt.Println(HumanReadableByteCount(452984832, false))
	fmt.Println(HumanReadableByteCount(28991029248, false))
	fmt.Println(HumanReadableByteCount(1855425871872, false))
	fmt.Println(HumanReadableByteCount(9223372036854775807, false))

}
