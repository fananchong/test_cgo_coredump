package main

import "C"
import "fmt"

//export OnExcept
func OnExcept(signum C.int) {
	fmt.Println("crash, signum:", signum)
}
