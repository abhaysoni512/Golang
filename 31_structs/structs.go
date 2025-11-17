package main

import (
	"fmt"
	"example.com/31_structs/user"
)



func main() {
	user_firstName := getUserData("Please enter your first name: ")
	user_lastName := getUserData("Please enter your last name: ")
	user_birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User = user.NewUser(user_firstName,user_lastName,user_birthdate)
	// ... do something awesome with that gathered data!
	
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	
	appUser.OutputUserDetails()
	
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}