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

type Admin struct {
	email    string
	password string
	User
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

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName:  "Admin",
			lastName:   "ADMIN",
			birthdate:  "----",
			createDate: time.Now(),
		},
	}

}

func (u User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
