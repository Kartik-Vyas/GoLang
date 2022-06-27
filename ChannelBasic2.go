package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() { //receiver goroutines
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() { //sender goroutines
			ch <- 42
			wg.Done()
		}()
	}

	wg.Wait()

}
