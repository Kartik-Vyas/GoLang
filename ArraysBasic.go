package main

import "fmt"

func main() {
	//declaring using index position
	grade := [3]int{97, 98, 99}
	fmt.Println(grade)

	//another way of declaring
	grade2 := [...]int{78, 86, 68, 97}
	fmt.Println(grade2)

	//using index position
	var student [3]string
	student[0] = "kartik"
	student[1] = "omkar"
	student[2] = "priyal"
	fmt.Println(student)

}
