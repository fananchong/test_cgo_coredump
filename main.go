package main

/*
extern void test_crash(char *str);
#cgo LDFLAGS: -L. -ltest
#include <stdlib.h>
*/
import "C"

import (
	_ "fmt"
	"unsafe"
)

func main() {
	str := "From Golang"

	//signal.Ignore(syscall.SIGABRT)

	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.test_crash(cStr)

	select {}
}
