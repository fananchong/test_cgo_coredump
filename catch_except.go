// +build !plan9,!windows

package main

/*
#include <signal.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

extern void onExcept(int signum);

static void handler(int signum) {
	char cmd[50];
	sprintf(cmd, "gcore %u", getpid());
	system(cmd);
	onExcept(signum);
	exit(1);
}

static void __attribute__ ((constructor)) sigsetup(void) {
	struct sigaction act;
	memset(&act, 0, sizeof act);
	act.sa_handler = handler;
	act.sa_flags = SA_RESETHAND;
	sigaction(SIGSEGV, &act, NULL);
	sigaction(SIGABRT, &act, NULL);
}
*/
import "C"
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

//export onExcept
func onExcept(signum C.int) {
	fmt.Println("crash, signum:", signum)

	// 打印 Crash 堆栈
	result := execCommand("./gdb_bt.sh", []string{"./main2", fmt.Sprintf("./core.%d", os.Getpid())})
	for _, line := range result {
		fmt.Printf(line)
	}
	fmt.Println("")
}

func sigsetup() {
	C.sigsetup()
}

func execCommand(commandName string, params []string) (result []string) {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Printf("执行命令: %s\n", strings.Join(cmd.Args, " "))
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error=>", err.Error())
		return nil
	}
	cmd.Start() // Start开始执行c包含的命令，但并不会等待该命令完成即返回。Wait方法会返回命令的返回状态码并在命令返回后释放相关的资源。

	reader := bufio.NewReader(pipe)

	var index int
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		result = append(result, line)
		index++
	}

	cmd.Wait()

	return result
}
