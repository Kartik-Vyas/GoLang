package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { //here since we have panic in other function it will terminate this once the panic happens but our main will exceute
			log.Println("Error:", err)
			panic(err) //sometime certain condition occur where we dont need to handle the error and want application to terminate,
					   // hence recalling panic which will terminate the application
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}
func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
