package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	createDate time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User

	appUser = User{
		firstName:  userFirstName,
		lastName:   userLastName,
		birthdate:  userBirthdate,
		createDate: time.Now(),
	}

	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func (u User) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}
