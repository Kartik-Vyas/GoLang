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
	m := map[[3]int]string{} //note we can have arrays as keys but we can't have slices as a key to a map
	fmt.Println(m)
	//other ways of declaring a map by using make function
	statePopulation2 := make(map[string]int)
	statePopulation2 = map[string]int{}
	fmt.Println(statePopulation2)
	//getting the values from maps
	fmt.Println(statePopulation["Ohio"])
	//adding value into the map
	statePopulation["Georgia"] = 1245255635
	fmt.Println(statePopulation)
	// Deleting the value in a map
	delete(statePopulation, "California")
	fmt.Println(statePopulation)
}
