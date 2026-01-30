package main

import "fmt"

// Type aliases for maps can be defined as types like so:
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// To optimize slices one can use "make" the define the type of the slice
	// and the initial length of the slice, since we know we're going to add
	// two items to this slice down below
	// In this case 5 is the max amount of items that can be added.
	userNames := make([]string, 2, 5)
	userNames[0] = "Jane"  // this is only allowed because we defined that the slice has two starting elements
	userNames[1] = "Frank" // if we tried doing this with a slice not created with make there would be an error

	// New arrays are created each time one appends to a slice
	// go manages these arrays behind the scenes and slices are
	// merely an abstraction.
	// if we use make on a slice and then append items will be appended to the values after the first two added by make
	userNames = append(userNames, "Dirk")
	userNames = append(userNames, "John")
	userNames = append(userNames, "Bill")
	fmt.Println(userNames)

	userNames = append(userNames, "Bob") // At this point the length of the slice defined when we called make will be exceeded and so go will make a new array behind the scenes
	fmt.Println(userNames)

	// When the map size needed is known in advance an optimization would be to use
	// make to specify that.
	// In this case we predefine the length of the map as having 3 items.
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	// courseRatings["vue"] = 4.9 // at this point the initial map size specified in the make function has been exceeded so go will have to create a new map behind the scenes
	// fmt.Println(courseRatings)
	courseRatings.output()

	// Looping over maps and slices/arrays using the "range" keyword
	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// maps can be iterated over in this case key will be of course the key!
	for key, value := range courseRatings {
		fmt.Printf("Key: %s, Value: %.1f\n", key, value)
	}
}
