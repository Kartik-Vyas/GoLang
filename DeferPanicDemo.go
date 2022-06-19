package main

import "fmt"

func main() {
	a := "start"
	defer fmt.Println(a) //defer check the value associated with variable at that time.
	a = "end"

	//panic
	// c, d := 1, 0
	// ans := c / d
	// fmt.Println(ans)
	//we can pass our own panic
	fmt.Println("start")
	defer fmt.Println("this was defered") //defer always exceute after main function and before panic
	panic("something went wrong")
	fmt.Println("end")
}
