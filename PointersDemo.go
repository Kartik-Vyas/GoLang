package main

import "fmt"

func main() {
	var a int = 42
	// b := a // since here b is copy only the value of a and not pointing towards it
	var b *int = &a    //state that b is pointer to an integer and it want to point to a
	fmt.Println(a, b)  //here b will store the memory location of a
	fmt.Println(a, *b) // using dereferancing opertor to pull the value
	a = 27             //now if we change the value of a
	fmt.Println(a, *b) // the value of b will now also chnage since it is pointing a.
	*b = 14
	fmt.Println(a, *b)
}
