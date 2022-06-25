package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Shows how many thread are available which is equal to number of core available in machine
	fmt.Println("Threads :", runtime.GOMAXPROCS(-1))
	//we can change the number of thread
	runtime.GOMAXPROCS(1)                            // truely concurrent application with no parrallelism
	fmt.Println("Threads :", runtime.GOMAXPROCS(-1)) //retuen the no. of thread previously set and -1 does change the value
}
