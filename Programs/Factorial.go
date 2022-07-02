package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number whose factorial you want..")
	fmt.Scan(&num)
	var fact int = 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	fmt.Println("the factorial is:", fact)
}
