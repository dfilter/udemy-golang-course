package main

import "fmt"

func main() {
	// Since two ints are passed in the infered type is int
	intResult := add(1, 2)
	fmt.Println("Int value:", intResult)

	// Since two strings are passed in the infered type is string
	stringResult := add("Hello, ", " World!")
	fmt.Println("Stromg value:", stringResult)

}

// generic T in this case can be one of int, float, or string. The same type is returned
func add[T int | float64 | string](a, b T) T {
	// The "+" operator is legal here because all the types T can be all support it
	return a + b
}
