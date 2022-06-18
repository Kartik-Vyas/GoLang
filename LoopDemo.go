package main

import (
	"fmt"
)

func main() {
	//we can have multiple declarations
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
Loop:
	for k := 0; k < 4; k++ {
		for j := 0; j < 7; j++ { //keeping loop till 7 for implementing break
			if j == 5 {
				break Loop // when we use break it get out of current loop but to implement break through out we declare it in following way.
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}
