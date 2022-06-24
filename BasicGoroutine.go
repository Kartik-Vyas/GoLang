package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()                      //Create a green Thread
	time.Sleep(100 * time.Millisecond) // Delays the main function so that it get time for for goroutine function to exceute.
}

//basic Goroutines
func sayHello() {
	fmt.Println("Hello")
}
