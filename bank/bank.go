package main

import (
	"fmt"

	"thiagoleite.me/go-bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 1000)

	showLogo()

	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println(("---------------------------------"))
		// panic("Can't continue, sorry :(")
	}

	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			printBalance(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using Go Bank")

			return
		}

		fmt.Println("=================================")
	}

}

func printBalance(balance float64) {
	fmt.Println("Your balance is: ", balance)
}
