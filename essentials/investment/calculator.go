package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years int
	var expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	// outputs information
	// fmt.Println("Future value:", futureValue)
	// fmt.Println("Future value (with inflation)", futureRealValue)

	// fmt.Printf("Future value: %.2f\n", futureValue)
	// fmt.Printf("Future value (with inflation): %.2f\n", futureRealValue)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future value (with inflation): %.2f\n", futureRealValue)

	formattedRFV := fmt.Sprintf(`\n
		Future value (with inflation): %.2f
		\n`, futureRealValue) // Using multiple lines

	fmt.Println(formattedFV)
	fmt.Println(formattedRFV)
}