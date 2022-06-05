package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Printf("Today is saturday")
	case today + 1:
		fmt.Printf("Tomorrow will be saturday")
	case today + 2:
		fmt.Printf("Will be in two days")
	default:
		fmt.Printf("Saturday is far away... ")
	}
}
