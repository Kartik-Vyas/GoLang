package main

import "fmt"

type Doctor struct { //capital letter for Doctor stating that it can be exported from this package
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number:    3,
		ActorName: "kartik",
		Companions: []string{
			"PV",
			"DV",
			"JV",
		},
	}
	// can be declared without using field name but it will cause confusion
	// eg aDoctor := Doctor {
	// 3,
	// "kartik",
	// []string{},
	//}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.ActorName)
	fmt.Println(aDoctor.Companions[1])
}
