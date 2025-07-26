package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	fileData, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}
	balaceText := string(fileData)
	balanceValue, err := strconv.ParseFloat(balaceText, 64)
	if err != nil {
		return 1000, errors.New("failed parse stored balance value")
	}
	return balanceValue, nil
}

func writeBalaceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Printf("ERROR \n%v\n-----\n", err)
		// to exit the process one can use the panic function
		// panic("You don't have an account with us you filthy hacker!")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Print(`What do you want to do?
1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
Your choice: `)

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Depost must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance:", accountBalance)
			writeBalaceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawl amount:")
			var withdrwalAmount float64
			fmt.Scan(&withdrwalAmount)

			if withdrwalAmount <= 0 {
				fmt.Println("Invalid amount. Withdrawl must be greater than zero.")
				continue
			}

			if withdrwalAmount > accountBalance {
				fmt.Println("Invalid amount. You don't have that much money in your account.")
				continue
			}

			accountBalance -= withdrwalAmount
			fmt.Println("Your new balance:", accountBalance)
			writeBalaceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go bank.")
			return
		}
	}
}
