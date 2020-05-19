package main

/*
extern void test_crash(char *str);
#cgo LDFLAGS: -L. -ltest
#cgo CFLAGS: -g3
*/
import "C"

import (
	_ "fmt"
)

func main() {
	str := "From Golang"

	//signal.Ignore(syscall.SIGABRT)

	cStr := C.CString(str)
	C.test_crash(cStr)

	select {}
}
