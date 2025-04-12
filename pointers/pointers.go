package main

import "fmt"

func main() {
	age := 32

	agePointer := &age

	fmt.Println("Age:", age)

	getAdultYears(agePointer)

	fmt.Println("Adult Years: ", age)

}

func getAdultYears(age *int) {
	*age -= 18
}
