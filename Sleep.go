package Goutils

import (
	"time"
)

func Sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
