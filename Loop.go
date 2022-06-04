package main

import "fmt"

func main() {

	// for {
	// 	fmt.Print("Hello world")
	// } Prints infinite time just like while

	var i int = 1
	for i <= 5 {
		fmt.Println("Hello world")
		i++
	}
	fmt.Println("-----------------")
	//
	for i = 1; i <= 5; i++ {
		fmt.Println("Hello world")
	}
}
