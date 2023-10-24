package Goutils

import "strconv"

// StringToInt converts a string to an integer.
func StringToInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}
