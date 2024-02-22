package main

import "fmt"

func PrintOptions() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func PrintLogo() {
	fmt.Println(`
	.d8888b.                888888b.                     888      
	d88P  Y88b               888  "88b                    888      
	888    888               888  .88P                    888      
	888         .d88b.       8888888K.   8888b.  88888b.  888  888 
	888  88888 d88""88b      888  "Y88b     "88b 888 "88b 888 .88P 
	888    888 888  888      888    888 .d888888 888  888 888888K  
	Y88b  d88P Y88..88P      888   d88P 888  888 888  888 888 "88b 
	 "Y8888P88  "Y88P"       8888888P"  "Y888888 888  888 888  888 																														 
	`)
}
