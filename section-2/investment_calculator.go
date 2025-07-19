package main

import (
	"fmt"
	"math"
)

func main() {
	// constants are of course immutable
	const inflationRate = 6.5
	// variables can be typed and mutated
	var investmentAmount float64
	var years float64
	// Variables can be defined and their type derived using ":="
	expectedReturnRate := 5.5

	// Print does not create a new line
	fmt.Print("Investment Amount: ")
	// "&" means that we pass a pointer to a variable to the function, meaning the variable will be mutated by the function
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Println prints the value to a new line
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
