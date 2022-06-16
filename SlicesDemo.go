package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a //In slices it directly refers to the arrays , hence any changes to b makes changes in a.
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("length is ", len(a))
	fmt.Println("Capacity is ", cap(a))
	//More on Slices
	p := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	q := p[:]   //slice of all elements
	r := p[3:]  // slice from 4th to end
	s := p[:6]  //slice first 6 elements
	t := p[3:6] //slice from 4th to 6th elements
	p[5] = 42   // changing the value here will still changes all the values since they are pointing to p.
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)

}
