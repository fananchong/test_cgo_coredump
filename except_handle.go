package main

import "C"
import "fmt"

//export onExcept
func onExcept(signum C.int) {
	fmt.Println("crash, signum:", signum)
}
