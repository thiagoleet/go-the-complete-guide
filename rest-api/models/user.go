package models

import (
	"errors"

	"thiagoleite.me/go-event-booking-api/db"
	"thiagoleite.me/go-event-booking-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	password, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, password)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswords(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
