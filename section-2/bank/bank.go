package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Printf(`What do you want to do?
1. Check balance
2. Deposit money
3. Withdraw money
4. Exit
Your choice: `)

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Depost must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance:", accountBalance)
		} else if choice == 3 {
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
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing Go bank.")
}
