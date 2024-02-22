package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)

	revenue = getUserInput("Enter revenue: ")

	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	expenses = getUserInput("Enter expenses: ")

	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)

	taxRate = getUserInput("Enter tax rate: ")

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateFinanctials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinanctials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}
