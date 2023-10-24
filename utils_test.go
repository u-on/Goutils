package Goutils

import "testing"

func TestRunCmd(t *testing.T) {
	err := RunCmd("ping", "www.bing.com")
	if err != nil {
		t.Log(err)

	} else {
		t.Log("success")
	}
}

func TestPathExists(t *testing.T) {
	exist, err := PathExists("D:\\iZOzone123.cf1")
	if err != nil {
		t.Log(err)

	} else {
		t.Log(exist)
	}
}

func TestReadFileToStr(t *testing.T) {
	str, err := ReadFileToStr("./ReadFileToStr.go")
	if err != nil {
		t.Log(err)

	} else {
		t.Log(str)
	}
}
