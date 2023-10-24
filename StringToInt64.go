package Goutils

import "strconv"

// StringToInt64 converts a string to an int64.
func StringToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
