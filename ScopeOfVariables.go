package main

import "fmt"

var a = 9

func demo() {
	var a = 8
	fmt.Println("In demo", a)
}
func main() {
	demo()
	fmt.Println("In Main", a)
}
