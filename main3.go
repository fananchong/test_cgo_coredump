package main

/*
#include "test.h"
#include <stdlib.h>
typedef void(*cb)(void);
#cgo LDFLAGS: -L. -lstdc++ -ltest
*/
import "C"

func main() {
	Sigsetup2()
	SafeCall(C.cb(C.test_crash2))
	select {}
}
