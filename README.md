# Go-shutdown
Go语言关机程序
package main：这一行声明了该文件属于 main 包，main 包是一个可执行程序的入口包。

import (...)：这里导入了两个包，fmt 用于格式化输出，os/exec 用于执行外部命令。

func main()：这是程序的入口点，所有的代码都在这里执行。

fmt.Println("关机程序启动！")：这一行使用 fmt.Println 函数在控制台输出一条消息。

cmd := exec.Command("shutdown", "/s", "/t", "0")：这一行创建了一个 exec.Command 对象，用于执行关机命令。"shutdown" 是要执行的命令，"/s" 和 "/t" 是命令的参数，"0" 是关机延迟时间。

err := cmd.Run()：这一行执行关机命令。cmd.Run() 方法会阻塞程序运行，直到关机命令执行完成。如果执行出现错误，将错误信息赋值给 err 变量。

if err != nil { fmt.Println("关机命令执行失败：", err) }：这一行检查关机命令是否执行失败。如果 err 不为 nil，即出现了错误，将错误信息输出到控制台。
