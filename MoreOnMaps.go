package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      65464354,
		"Florida":    455624562,
		"Ohio":       458752156,
	}
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["oho"]) //since we do not have the key with this name so it give the 0 output
	//now to know whether our map contains that key we can do following code
	pop, ok := statePopulation["oho"]
	fmt.Println(pop, ok) // it will five o as an output also with boolean value
	// to only check for a key
	_, okay := statePopulation["Ohio"]
	fmt.Println(okay) //generally we use ok for checking but using okay since it is already declared.
	//for geeting the length
	fmt.Println("length is ", len(statePopulation))
	//same as slices it copies the reference of map hence changing value in there also affects the current map
	sp := statePopulation
	delete(sp, "Florida")
	fmt.Println(statePopulation)
}
