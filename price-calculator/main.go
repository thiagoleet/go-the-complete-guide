package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		const intputFilePath = "prices.txt"
		outputPath := fmt.Sprintf("result_%.0f.json", taxRate*100)

		fm := filemanager.New(intputFilePath, outputPath)

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}
