package main

import "fmt"

func say(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("the value of the index is:", idx)
}
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("the sum is ", result)
}
func main() {
	for i := 0; i < 5; i++ {
		say("Hello Go!", i)
	}
	sum(1, 2, 3, 4, 5)
}
