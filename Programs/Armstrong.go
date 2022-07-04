package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter any 3 digit number")
	var num int
	fmt.Scan(&num)
	var temp int = num
	var sum int
	var val int
	for num != 0 {
		val = num % 10
		val = int(math.Pow(float64(val), 3))
		num = num / 10
		sum += val
	}
	if sum == temp {
		fmt.Println("The number is an Armstrong number")
	} else {
		fmt.Println("The number is not an Armstrong number")
	}
}
