package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Thiago")
	userNames = append(userNames, "John")

	fmt.Println(userNames, len(userNames), cap(userNames))

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9
	courseRatings["node"] = 4.5

	fmt.Println(courseRatings, len(courseRatings))
}
