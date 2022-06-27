package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) { //Only receiving the data
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) { //only sending the data
		i := 42
		ch <- i
		wg.Done()
	}(ch)
	wg.Wait()
}
