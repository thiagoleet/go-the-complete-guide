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

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
