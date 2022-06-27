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
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
  	ch <- i //channel here copies same as other variable but not the referance
		i = 27  //hence the value do not change when re-assign
		wg.Done()
	}()
	wg.Wait()
}
