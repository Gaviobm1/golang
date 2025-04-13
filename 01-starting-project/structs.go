package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	admin := user.NewAdmin("example@email.com", "password1")

	admin.OutputUserData()

	user, error := user.New(firstName, lastName, birthDate)
	// ... do something awesome with that gathered data!
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	user.OutputUserData()
	user.ClearUserName()
	user.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
