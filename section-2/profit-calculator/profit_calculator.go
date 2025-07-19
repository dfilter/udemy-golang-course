package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expense)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expense
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Print("Earnings Before Tax: ")
	fmt.Println(earningsBeforeTax)

	fmt.Print("Profit: ")
	fmt.Println(profit)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}
