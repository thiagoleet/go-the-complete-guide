package main

import "fmt"

func main() {
	var productNames [4]string
	productNames = [4]string{"A Book"}

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[0])

	// Slices
	featuredPrice := prices[1:3]
	fmt.Println(featuredPrice)
	fmt.Println(len(featuredPrice), cap(featuredPrice))

	featuredPrice = prices[:3]
	fmt.Println(featuredPrice)
	fmt.Println(len(featuredPrice), cap(featuredPrice))

	featuredPrice = prices[1:]
	fmt.Println(len(featuredPrice), cap(featuredPrice))

}
