// +build !plan9,!windows

package main

/*
#include "breakpad.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -L./breakpad/lib/ -lstdc++ -lbreakpad_client -ltest
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export onDumpCallback
func onDumpCallback() {
	fmt.Println("call onDumpCallback")
}

func sigsetup() {
	str := "."
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.breakpad_init(cStr)
}
