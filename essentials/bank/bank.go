package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)

	balance, _ := strconv.ParseFloat(balanceText, 64)

	return balance
}

func main() {
	var accountBalance float64 = getBalanceFromFile()

	PrintLogo()

	for {
		PrintOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		// if choice == 1 {
		// 	printBalance(accountBalance)
		// } else if choice == 2 {
		// 	fmt.Print("Your deposit: ")
		// 	var depositAmmount float64
		// 	fmt.Scan(&depositAmmount)

		// 	if depositAmmount < 0 {
		// 		fmt.Println("Invalid amount! Must be greater than zero")
		// 		// return // stop the execution of the function
		// 		continue
		// 	}
		// 	accountBalance += depositAmmount

		// 	printBalance(accountBalance)
		// } else if choice == 3 {
		// 	fmt.Print("Withdraw amount: ")
		// 	var withdrawAmmount float64
		// 	fmt.Scan(&withdrawAmmount)

		// 	if withdrawAmmount < 0 {
		// 		fmt.Println("Invalid amount! Must be greater than zero")
		// 		continue
		// 	} else if withdrawAmmount > accountBalance {
		// 		fmt.Println("Invalid amount! You can't withdraw more than you have!")
		// 		continue
		// 	}

		// 	accountBalance -= withdrawAmmount

		// 	printBalance(accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }

		switch choice {
		case 1:
			printBalance(accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmmount float64
			fmt.Scan(&depositAmmount)

			if depositAmmount < 0 {
				fmt.Println("Invalid amount! Must be greater than zero")
				// return // stop the execution of the function
				continue
			}
			accountBalance += depositAmmount
			writeBalanceToFile(accountBalance)
			printBalance(accountBalance)
		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmmount float64
			fmt.Scan(&withdrawAmmount)

			if withdrawAmmount < 0 {
				fmt.Println("Invalid amount! Must be greater than zero")
				continue
			} else if withdrawAmmount > accountBalance {
				fmt.Println("Invalid amount! You can't withdraw more than you have!")
				continue
			}

			accountBalance -= withdrawAmmount
			writeBalanceToFile(accountBalance)
			printBalance(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank")

			return
		}
	}

}

func printBalance(balance float64) {
	fmt.Println("Your balance is: ", balance)
}
