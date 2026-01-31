package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	if number == 1 {
		return number
	}

	// This recursive call will continue untill we reach the point where 1 is returend above (added to the call stack)
	// Once it reaches 1 factorial will reach the exit case and not call its self recursively
	// factorial will return 1
	// and then calcuate 1 * 2 and return 2
	// then 2 * 3 and return 6
	// then 6 * 4 and return 24
	// then 24 * 5 and finally return 120
	return number * factorial(number-1)
}
