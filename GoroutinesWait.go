package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"     //Here the variable declare in main function can be access by the underlying function
	wg.Add(1)             //Synchronize the goroutine , here we have only one hence 1 in argument.
	go func(msg string) { //Since func can take arguments so here it copies the values of msg when we were innvoking the func of goroutine
		fmt.Println(msg)
		wg.Done() // notify the wait method that we are completed with this function.
		//it will decrement the no. of group by 1 hence here we have one goroutine hence decremented by one and wait method will finish the application
	}(msg)
	msg = "Goodbye" // the main thread is still running before the sleep so it reassign the msg variable before the goroutine is exceuted
	wg.Wait()       //will wait till function is done
	//time.Sleep(100 * time.Millisecond) Here we are not rely on any timer so exceution becomes fast
}
