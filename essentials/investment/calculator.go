package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years int
	var expectedReturnRate float64

	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, float64(years))

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

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) {

	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv := fv / math.Pow(1+inflationRate/100, float64(years))

	return fv, rfv
}
