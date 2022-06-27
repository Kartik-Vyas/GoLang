package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch      // reading into channel
		fmt.Println(i) //printing value
		ch <- 27       //again writing into channel
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i           //writing into channel
		fmt.Println(<-ch) //printing the value writen into channel
		wg.Done()
	}()
	wg.Wait()
}
