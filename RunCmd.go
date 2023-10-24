package Goutils

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"runtime"
)

func RunCmd(command string, args ...string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)

	// 获取命令的标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error obtaining stdout pipe:", err)
		return err
	}

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return err
	}

	// 实时读取并处理命令输出
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if runtime.GOOS == "windows" { //处理windows中文乱码
			fmt.Println(Utf8ToGbk(line))
		} else {
			fmt.Println(line)
		}

		// 可以在这里添加更多的处理逻辑，如实时记录日志或实时显示进度等
	}

	// 等待命令完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command to finish:", err)
		return err
	}

	//fmt.Println("Command execution completed.")
	return nil
}
