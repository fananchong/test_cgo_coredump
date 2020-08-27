package main

/*
#include "test.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -ltest
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
	// C.test_crash2()
	C.test_crash(cStr)

	select {}
}
