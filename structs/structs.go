package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	createDate time.Time
}

func newUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &User{
		firstName:  firstName,
		lastName:   lastName,
		birthdate:  birthdate,
		createDate: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *User
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return

	}

	appUser.outputUserData()
	appUser.clearUserName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func (u User) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
