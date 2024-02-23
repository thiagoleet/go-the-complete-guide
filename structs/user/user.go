package user

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

func New(firstName, lastName, birthdate string) (*User, error) {
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

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
