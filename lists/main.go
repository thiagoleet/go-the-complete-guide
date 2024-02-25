package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Thiago")
	userNames = append(userNames, "John")

	fmt.Println(userNames, len(userNames), cap(userNames))
}
