package main

import "fmt"

func main() {
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")
	fmt.Println()
	//using defer
	fmt.Println("start")
	defer fmt.Println("middle") //defer exceutes after the main func exceution is completed
	fmt.Println("end")
	fmt.Println()

	//having multiple defer follow lifo order
	defer fmt.Println("begin")
	defer fmt.Println("center") //defer exceutes after the main func exceution is completed
	defer fmt.Println("stop")
}
