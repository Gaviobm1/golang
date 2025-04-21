package models

import "example.com/restful/db"

type User struct {
	ID              int64
	Email, Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO events(email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}
