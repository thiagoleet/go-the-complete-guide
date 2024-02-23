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

func newUser(firstName, lastName, birthdate string) User {
	return User{
		firstName:  firstName,
		lastName:   lastName,
		birthdate:  birthdate,
		createDate: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User = newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputUserData()
	appUser.clearUserName()
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

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
