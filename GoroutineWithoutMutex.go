package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func sayHello() {
	fmt.Println("Hello", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}
