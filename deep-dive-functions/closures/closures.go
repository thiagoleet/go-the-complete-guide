package closures

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	double := createTramsformer(2)
	triple := createTramsformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTramsformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
