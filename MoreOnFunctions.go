package main

import "fmt"

func sum(value ...int) *int { //there is a pointer to an integer pointing to result
	fmt.Println(value)
	result := 0
	for _, v := range value {
		result += v
	}
	return &result //returning the address of result
}
func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println(*s) //de-referancing since it is returning pointer
}
