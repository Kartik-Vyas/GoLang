package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number whose sum you want till that number")
	fmt.Scan(&num)
	var sum int = 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Println("The sum of number is:", sum)
}
