package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile("balance.txt", []byte(balanceText), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Balance saved to file.")
	}
}

func main() {
	var accountBalance float64 = 10000

	fmt.Println("Welcome to Go Bank Vsense!")

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit ")

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Enter deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance after deposit is:", accountBalance)
			writeBalanceToFile(accountBalance)

		case 3:
			fmt.Print("Enter withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New balance after withdrawal is:", accountBalance)
			writeBalanceToFile(accountBalance)

		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
