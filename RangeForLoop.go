package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for k, v := range s { //here key are index position and values are real value
		fmt.Println(k, v)
	}

	sen := "hello Go!"
	for k, v := range sen { // here values are displayed on there asci value
		fmt.Println(k, string(v))
	}

	statePopulation := map[string]int{
		"California": 39250017,
		"Texas":      65464354,
		"Florida":    455624562,
		"Ohio":       458752156,
	}
	for k, v := range statePopulation { //we can have single key without values but for only values we need to use
		// for _,v := range statepopulation{
		//fmt.Println(v)
		//}
		fmt.Println(k, v)
	}
}
