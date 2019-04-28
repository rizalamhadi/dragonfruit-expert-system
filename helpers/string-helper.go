package helpers

import (
	"strconv"
)

func StringToInt(s string) int {
	u64, _ := strconv.ParseUint(s, 10, 32)

	return int(u64)
}