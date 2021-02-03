package main

/*
#include "test.h"
#include "catch_except.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -ltest
*/
import "C"

import (
	_ "fmt"
	"unsafe"
)

func main() {
	C.sigsetup()

	str := "From Golang"
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	// C.test_crash2()
	C.test_crash(cStr)
	// testCrash3()

	select {}
}
