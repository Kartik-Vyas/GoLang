package main

import "fmt"

func mars_age(age int) int {
	age = age * 365
	mage := age / 687
	return mage

}

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}
