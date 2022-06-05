package main

import "fmt"

func main() {
	var num = 2
	//else statement to be used just besides the closing bracket of if statement
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Your number is greater")
	}
}
