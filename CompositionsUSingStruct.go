package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max : "100"` // used in case of web application where it specifies name is req and max character must be 100
	Origin string
}

type Bird struct {
	Animal   // embeding the struct of animal into the bird
	SpeedKPH float32
	Canfly   bool
}

// This is not same as inheritance since bird is still an independent struct it is only embeding the animal struct, hence having HAS-A relationship
func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Austrialia"
	b.SpeedKPH = 42
	b.Canfly = false
	fmt.Println(b)
	fmt.Println(b.Name) // can also get the single entity
	//can use explicitly using animal struct called literal syntax
	c := Bird{
		Animal:   Animal{Name: "Rushky", Origin: "Russia"},
		SpeedKPH: 30,
		Canfly:   false,
	}
	fmt.Println(c)

	//Tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
