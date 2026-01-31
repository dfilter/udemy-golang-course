package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Here we use an anonymous function instead of passing a reference to a predefined function
	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformedNumbers)

	double := createTransformer(2)
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println(doubledNumbers)

	tripple := createTransformer(3)
	trippledNumbers := transformNumbers(&numbers, tripple)
	fmt.Println(trippledNumbers)

	counter1 := counter()
	fmt.Println(counter1()) // will print 1
	fmt.Println(counter1()) // will print 2
	fmt.Println(counter1()) // with print 3

	counter2 := counter()
	fmt.Println(counter2()) // will print 1 and not 4
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := make([]int, len(*numbers))
	for i, value := range *numbers {
		dNumbers[i] = transform(value)
	}
	return dNumbers
}

// Factory function design pattern; a function that create other functions
func createTransformer(factor int) func(int) int {
	// Here factor is baked into the returned function since its provided from the outer scope making this a closure
	return func(number int) int {
		return number * factor
	}
}

// This function exemplifies a closure well because each time the returned function is called count will be incremented
// If a different instance of counter is created the count variable will start at 0 again
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
