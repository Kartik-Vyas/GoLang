package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c)

	var ms *myStruct
	fmt.Println(ms)
	ms = &myStruct{ // ms is holding the address of an object that has a field holding value 42
		foo: 42,
	}
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
