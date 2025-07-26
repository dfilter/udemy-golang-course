package user

import (
	"errors"
	"fmt"
	"time"
)

// Public/private access doesn't just relate to the struct name but also its fields
// therefor public fields need to be capitalized if they need to be publically acessible
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	// This is called explicit embedding and provides a field to access the embedded User struct
	// User     User

	// anonymous embedding merges the fields of the User struct with this struct
	User
}

// adding "(u User)" to this function makes it a "method" of the "User" struct
// in this case "u" is the instance of the struct
// "(u User)" is called a receiver argument
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// since the data in the struct is being mutated the reference needs to be used instead "(u *User)""
// so that its not just a copy of the struct
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}
