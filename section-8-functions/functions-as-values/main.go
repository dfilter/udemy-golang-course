package main

import "fmt"

/*
functions are first class values and can be passed as values to other functions
*/

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	moreNumbers := []int{5, 1, 2}

	// here we are passing a function reference to transformNumbers just like any other value
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	trippled := transformNumbers(&numbers, triple)
	fmt.Println(trippled)

	// This value is going to be a reference to a function
	transformerFn1 := getTransformFunction(&numbers)
	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	fmt.Println(transformedNumbers)

	transformerFn2 := getTransformFunction(&moreNumbers)
	transformedNumbers2 := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers2)
}

// This type constrains to anything where the underlying type is numeric
type Number interface {
	~int | ~int64 | ~float32 | ~float64 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~uintptr | ~complex64 | ~complex128
}

// This function type takes a generic type of Number and returns that same type
type NumberFunction[T Number] func(T) T

// Here we define a function as a parameter of another number
func transformNumbers[T Number](numbers *[]T, transform NumberFunction[T]) []T {
	dNumbers := make([]T, len(*numbers))
	for i, value := range *numbers {
		dNumbers[i] = transform(value)
	}
	return dNumbers
}

// Get transformFunction returns a reference to double function or triple function
func getTransformFunction[T Number](numbers *[]T) NumberFunction[T] {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double[T Number](number T) T {
	return number * 2
}

func triple[T Number](number T) T {
	return number * 3
}
