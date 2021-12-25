package main

/*
#include "test.h"
#include <stdlib.h>
typedef void(*cb)(void);
#cgo LDFLAGS: -L. -lstdc++ -ltest
*/
import "C"

import (
	_ "fmt"
)

func main() {
	Sigsetup2()

	SafeCall(C.cb(C.test_crash2))

	select {}
}
