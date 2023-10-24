package Goutils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
)

// RunCmd2 运行给定的命令及参数。
//
// 它接受一个命令字符串和可变参数作为额外的命令参数。
// 如果执行命令时出现问题，返回一个错误。
func RunCmd2(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	// 创建一个管道来捕获命令的输出
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", err)
		return err
	}

	// 创建一个扫描器来读取命令输出的内容
	scanner := bufio.NewScanner(cmdReader)

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting Cmd", err)
		return err
	}

	// 创建一个信号通道来捕获Ctrl+z或Ctrl+C
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	// 使用goroutine来持续读取命令输出并显示进度
	go func() {
		for scanner.Scan() {
			if runtime.GOOS == "windows" { //处理windows中文乱码
				fmt.Println(Utf8ToGbk(scanner.Text()))

			} else {
				fmt.Println(scanner.Text())
			}

		}
	}()

	// 等待Ctrl+z或Ctrl+C的信号
	<-signalChannel

	// 终止进程
	err = cmd.Process.Kill()
	if err != nil {
		//fmt.Println("Error killing Cmd", err)
		return err
	}

	// 等待命令完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for Cmd to finish", err)
		return err
	}
	return nil
}
