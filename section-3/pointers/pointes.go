package main

import "fmt"

// "*int" means that age is a reference to an iteger
func getAdultYears(age *int) {
	// to do the calculation age since its a reference needs to be dereferenced (value needs to be looked up)
	// return *age - 18
	// to directly alter the value in memory we can dereference age
	*age -= 18
}

func main() {
	age := 30 // regular variable

	var agePointer *int
	agePointer = &age
	// To dereference a pointer (use its value) one can put "*" in front of it.
	fmt.Println("Age:", *agePointer)
	// fmt.Println("Age Pointer address:", agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult years:", age)
}
