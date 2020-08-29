package main

/*
#include "test.h"
#include <stdlib.h>
#cgo LDFLAGS: -L. -ltest
*/
import "C"

import (
	"fmt"
	_ "fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	str := "From Golang"
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	// C.test_crash2()
	C.test_crash(cStr)
	// testCrash3()

	select {}
}

func testCrash3() {

	go func() {

		defer func() {
			fmt.Println("exit goroutine#2")
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		time.Sleep(2 * time.Second)
		str := "From Golang"
		cStr := C.CString(str)
		defer C.free(unsafe.Pointer(cStr))
		C.test_crash(cStr)
		fmt.Println("exit goroutine#1")
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGABRT)
	for {
		fmt.Println("xxxxxxxxxxxxxxxxxxxx")
		select {
		case s := <-sig:
			fmt.Printf("receive signal : %v\n", s)
		}
	}
}
