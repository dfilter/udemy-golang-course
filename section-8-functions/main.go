package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	// Numbers can be pulled out of their slice by using ... turning each value into its own arugment
	sumSlice := sumup(1, numbers...)
	fmt.Println(sumSlice)

	sum := sumup(1, 10, 15)
	fmt.Println(sum)
}

// This is a variadic function in that it can take any number of aguments in this case of type int
func sumup(startingNumber int, numbers ...int) int {
	sum := startingNumber

	for _, val := range numbers {
		sum += val
	}

	return sum
}
