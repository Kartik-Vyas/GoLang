package main

import (
	"fmt"
	"sync"
)

//by using wait method we are not following the parallelism because thread will wait for other thread to complete the execution
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHello() {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}
