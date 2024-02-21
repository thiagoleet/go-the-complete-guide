package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years = 1000.0, 10
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), float64(years))

	fmt.Println(futureValue)

}
