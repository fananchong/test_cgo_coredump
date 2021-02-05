package main

/*
#include "test.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -L./breakpad/lib/ -lstdc++ -lbreakpad_client -ltest
*/
import "C"

import (
	"fmt"
	_ "fmt"
	"time"
	"unsafe"
)

func main() {
	Sigsetup()

	gopanic()
	gopanic()

	fmt.Println("1")

	time.Sleep(1 * time.Second)
	str := "From Golang"
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	C.test_crash2()
	C.test_crash2()
	// C.test_crash(cStr)
	// testCrash3()

	select {}
}

func gopanic() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	var a *int
	fmt.Println(*a)
}
