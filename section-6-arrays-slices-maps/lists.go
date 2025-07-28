package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices[1])
	prices[1] = 9.99
	// prices[2] = 11.99 // this will throw an error since the slice was created with only two items.
	// so instead one can use:
	updatedPrices := append(prices, 11.99)
	fmt.Printf("%v <- Original slice did not update\n", prices)
	fmt.Printf("%v <- New slice has the appened value\n", updatedPrices)
}

func arraysAndSlices() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)
	fmt.Println(prices[2])

	// Length of an array needs to be known at compile time and can't be assigned dynamically
	// Define an array of type string with 4 indexes 0,1,2,3
	var productNames [4]string = [4]string{"A book", "A thing", "stuff", "another thing"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames[2])

	// featuredPrices is whats called a "slice" its an array with no set length
	// like in python colon can be used to get specific portions (slice) of an array
	featuredPrices := prices[1:] // everything but the first item

	// like in most other languages go arrays and slices are reference types, meaning if you modify a
	// a value in a slice of an array it'll modify the value in the original
	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices)
	// len() builtin function that gets the length of an array/slice
	fmt.Println(len(featuredPrices))
	// cap() builtin function that gets the capacity of an array/slice
	// in this case highlightedPrices is a slice based on the "prices" array excluding the first element
	// that element is lost for good after slicing but values the the right are still accessible and therefor cap() results in 3 despite there only being 1 element in the slice
	highlightedPrices := featuredPrices[:1]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	featuredPrices2 := prices[1:3] // everything from index 1 up to index 3
	fmt.Println(featuredPrices2)
	featuredPrices3 := prices[:4] // everything but the last item
	fmt.Println(featuredPrices3)
}
