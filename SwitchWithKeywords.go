package main

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one five or ten")
	case 2, 4, 6:
		fmt.Println("two four or six")
	default:
		fmt.Println("not in the following case")
	}

	num := 10
	switch {
	case num <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // note fallthrough will mke the code exceute the case below it without check the condition
	case num <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("number greater than 20")
	}

	//type Switch
	var numb interface{} = []int{}
	switch numb.(type) {
	case int:
		fmt.Println("number is an int")
	case float64:
		fmt.Println("number is a float64")
	case string:
		fmt.Println("number is a string")
	case []int:
		fmt.Println("number is a slice")
		break //this will not allow to exceute further in this case
		fmt.Println("break this statement")
	default:
		fmt.Println("unmatched type")
	}
}
