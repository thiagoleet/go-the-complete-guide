package main

import "fmt"

func main() {
	age := 32
	var adultYears int

	var agePointer *int
	agePointer = &age

	// fmt.Println("Age:", age)
	fmt.Println("Age:", *agePointer)
	// fmt.Println(agePointer)

	setAdultYears(agePointer, &adultYears)
	fmt.Println(adultYears)
}

func setAdultYears(age, adultYears *int) {
	*adultYears = *age - 18
}
