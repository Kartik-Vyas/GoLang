package main

import "fmt"

func main() {
	make function
	k := make([]int, 3, 100)
	fmt.Println(k)
	fmt.Println("length is ", len(k))
	fmt.Println("Capacity is ", cap(k))

	a := []int{}
	fmt.Println(a)
	fmt.Println("length is ", len(a))
	fmt.Println("Capacity is ", cap(a))

	a = append(a, 1)
	fmt.Println(a)
	fmt.Println("length is ", len(a))
	fmt.Println("Capacity is ", cap(a))

	a = append(a, 2, 3, 4, 5) // can be also done as shown below both are same
	//a = apppend(a, []int{2,3,4,5}....) this called spread opertaor
	fmt.Println(a)
	fmt.Println("length is ", len(a))
	fmt.Println("Capacity is ", cap(a))

	//removing elements from slices
	//Note since we are working with slices hence it will affect the unlining array hence result may vary.
	b := a[1:]                   //pops the first element
	c := a[:len(a)-1]            // pops the last element
	d := append(a[:2], a[3:]...) // pops the middle element , note it will completely remove three from underlying array of a
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
