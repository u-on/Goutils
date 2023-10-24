package Goutils

import "strconv"

// Int64ToString converts an int64 to a string.
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
