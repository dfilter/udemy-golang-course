package main

import "fmt"

/*
functions are first class values and can be passed as values to other functions
*/

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)

}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, double(value))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
