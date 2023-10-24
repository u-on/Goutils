package Goutils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Pause() {
	var s string
	print("Press enter to continue...")
	_, _ = fmt.Scanln(&s)
}

// PauseExit Press Ctrl+c to exit
func PauseExit() {
	fmt.Printf("Press Ctrl+c to exit...")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	for {
		if sig.String() == "interrupt" {
			break
		}
	}
	return

}
