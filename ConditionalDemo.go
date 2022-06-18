package main

import (
	"fmt"
)

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}

func main() {
	num := 50
	guess := 30
	if guess < 1 || returnTrue() || guess > 100 { // shortcircuiting takes place whenever the condition is true and it will not exceute further statement of ||
		fmt.Println("The guess num must be within the range")
	}
	if guess < 1 {
		fmt.Println("number less than one")
	} else if guess > 100 {
		fmt.Println("number more than 100")
	} else {
		fmt.Println(num <= guess, num >= guess, num != guess)
	}
	//if guess >= 1 && guess <= 100 {
	if guess < num {
		fmt.Println("too low")
	} else if guess > num {
		fmt.Println("too high")
	} else if guess == num {
		fmt.Println("correct")
	}

}
