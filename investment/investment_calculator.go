package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64 = 1000.0
	var years = 10
	expectedReturnRate := 5.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
