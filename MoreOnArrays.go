package main

import "fmt"

func main() {
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)
	// other ways for 2d matrix
	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[0] = [3]int{0, 1, 0}
	identityMatrix2[0] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)
	//more on arrays
	a := [...]int{1, 2, 3}
	b := a   // note it store the copy of an array a
	b[1] = 5 // Hence while change the first element of b is not same as elements of a since it is a copy of a.
	c := &a
	c[2] = 4 // Here we are directly pointing to the a hence if we update the element of c it will affect the elements of a.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
