package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("关机程序启动！")
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println("关机命令执行失败：", err)
	}
}
