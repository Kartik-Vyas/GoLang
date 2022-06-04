package main

import "fmt"

func add(num1, num2 int) int {
	var out = num1 + num2
	return out
}
func calc(num1, num2 int) (int, int) {
	var addition = num1 + num2
	var subtraction = num1 - num2
	return addition, subtraction
}
func main() {
	num1 := 7
	num2 := 5
	//calling the function
	result := add(num1, num2)
	//go lang can have multiple return statement
	result1, result2 := calc(num1, num2)
	fmt.Println(result)
	fmt.Println("The first value is:", result1, "and Second value is:", result2)
}
