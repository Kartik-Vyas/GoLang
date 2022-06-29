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
		for {
			if i, ok := <-ch; ok { //the values enter the variable and return boolean value for ok
				fmt.Println(i) //if the value dont enter than it return false hence else block exceutes
			} else {
				break
			} // similar exceution ocurr in case of for range loop
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
