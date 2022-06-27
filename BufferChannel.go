package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) //created a buffer that can store 50 integers
	// But here we lost our data i.e 27 which is entering into channel
	wg.Add(2)
	go func(ch <-chan int) { //Only receiving the data
		i := <-ch
		fmt.Println(i)
		i = <-ch //creating one more channel receiver to eliminate data loss
		fmt.Println(i)
		wg.Done()
	}(ch) // polymorphic behaver since passing receiver full type but here defining only ch which is understand by Go
	go func(ch chan<- int) { //only sending the data
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}
