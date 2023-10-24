package Goutils

import "os"

func IsDir(path string) bool {
	s, _ := os.Stat(path)
	return s != nil && s.IsDir()

}
