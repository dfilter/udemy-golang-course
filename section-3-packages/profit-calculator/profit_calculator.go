package main

import (
	"fmt"
	"os"
)

const fileName = "results.txt"

func storeResults(formattedResults string) {
	os.WriteFile(fileName, []byte(formattedResults), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (earningsBeforeTax, profit, ratio float64) {
	earningsBeforeTax = revenue - expenses
	profit = earningsBeforeTax * (1 - taxRate/100)
	ratio = earningsBeforeTax / profit
	return
}

func getUserInput(message string) (float64, error) {
	fmt.Print(message)
	var userInput float64
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, fmt.Errorf("invalid input %v must be greater than zero", message)
	}
	return userInput, nil
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expense, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	formattedResults := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.3f\n", earningsBeforeTax, profit, ratio)
	storeResults(formattedResults)

	fmt.Print(formattedResults)
}
