package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName, lastName, birthDate string
	createdAt                      time.Time
}

type Admin struct {
	email, password string
	User
}

func (u *User) OutputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Este",
			birthDate: "----",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" || lastName == "" {
		return nil, errors.New("First/Last name cannot be blank")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
