package Goutils

import "os"

// WriteStringToFile writes a string to a file.
//
// str is the string to be written to the file.
// fileName is the name of the file to write to.
// Returns an error if there was an issue writing the file.
func WriteStringToFile(str string, fileName string) error {
	return os.WriteFile(fileName, []byte(str), 0666)
}
