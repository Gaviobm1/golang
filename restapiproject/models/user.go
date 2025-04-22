package models

import (
	"errors"

	"example.com/restful/db"
	"example.com/restful/utils"
)

type User struct {
	ID              int64
	Email, Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	isValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !isValid {
		return errors.New("credentials invalid")
	}

	return nil
}
