package main

import "fmt"

func main() {
	fmt.Println("Enter the three numbers")
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c int
	fmt.Scan(&c)
	if a > b && a > c {
		fmt.Println("The Greatest number is:", a)
	} else if b > a && b > c {
		fmt.Println("The Greatest number is:", b)
	} else if c > a && c > b {
		fmt.Println("The Greatest number is:", c)
	} else {
		fmt.Println("Invalid Input")
	}
}
