package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// execCommand("echo", []string{"xxx"})
	execCommand("./main", []string{})
	execCommand("zip", []string{"core.zip", "./core"}) // 这里仅为例子，也可以调用 shell 脚本
}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	cmd.Env = append(cmd.Env, "GOTRACEBACK=crash")
	//显示运行的命令
	fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args, " "))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return false
	}
	cmd.Start() // Start开始执行c包含的命令，但并不会等待该命令完成即返回。Wait方法会返回命令的返回状态码并在命令返回后释放相关的资源。

	reader := bufio.NewReader(stdout)

	var index int
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
		index++
	}

	cmd.Wait()

	return true
}
