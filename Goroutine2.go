package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello" //Here the variable declare in main function can be access by the underlying function
	// go func() {       //Anonymous function
  // 	fmt.Println(msg) // output using this would be (Goodbye)
	// }()
	go func(msg string) { //Since func can take arguments so here it copies the values of msg when we were innvoking the func of goroutine
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye" // the main thread is still running before the sleep so it reassign the msg variable before the goroutine is exceuted
	time.Sleep(100 * time.Millisecond)
}
