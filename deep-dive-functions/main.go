package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	sum := sumUp(1, 10, 15, 40, -5)

	fmt.Println(sum)

	anotherSum := sumUp(1, numbers...)
	fmt.Println(anotherSum)
}

func sumUp(starting int, numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
