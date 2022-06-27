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
		for i := range ch { //for receiving multiple data from channel we are using range of channel
			fmt.Println(i) // since range in waiting for receiving the elements but there are no more than 2 values hence it gives error
		}
		wg.Done()
	}(ch) // polymorphic behaver since passing receiver full type but here defining only ch which is understand by Go
	go func(ch chan<- int) { //only sending the data
		ch <- 42
		ch <- 27
		close(ch) // close function tells the channel that we are done sending the data
		// note:= closing channel is irreversible
		wg.Done()
	}(ch)
	wg.Wait()
}
