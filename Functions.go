package main

import "fmt"

func add(num1, num2 int) int {
	var out = num1 + num2
	return out
}
func main() {
	num1 := 5
	num2 := 7
	//calling the function
	result := add(num1, num2)
	fmt.Println(result)
}
