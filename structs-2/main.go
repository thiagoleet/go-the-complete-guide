package main

import "fmt"

// creating an alias to string
type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "John"

	name.log()

}
