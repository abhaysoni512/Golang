package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName ,lastName ,birthdate string
	createdAt time.Time
}
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName(){
	u.firstName = ""
	u.lastName = ""
}


func NewUser(firstName, lastName, birthdate string) *User {
	if firstName == "" ||  lastName == "" || birthdate == "" {
		panic("All fields are required to create a new User.")
	}
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}