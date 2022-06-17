package main

import "fmt"

func main() {
	//creating anonymous struct
	aDoctor := struct{ name string }{name: "kartik"} //  it does not have a name hence can be used only for limited purpose.
	fmt.Println(aDoctor)
	// struct does not point to the reference like slices it create the copy hence do not affect the original struct
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom"
	fmt.Println(anotherDoctor)
	fmt.Println(aDoctor)
	doc := &aDoctor // just like arrays we can use address operator
	fmt.Println(doc)
}
