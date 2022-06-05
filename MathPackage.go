package main

import (
	"fmt"
	"math" //For all the imports done the function name must start from capital as we are importing outside from that package
)

//But for main since we are creating it hence it is declared with small letter
func main() {
	var num float64 = 12
	var result float64 = math.Sqrt(num)
	fmt.Println("The actual value is ", result)
	//Other math functions
	fmt.Printf("This will get only two digits after decimal %.2f", result)
	fmt.Println()
	fmt.Println("This will get the ceiling value", math.Ceil(result))
	fmt.Println("This will get the floor value", math.Floor(result))

}
