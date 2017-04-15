package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User - model
// TODO - add hashing function | don't store plaintext password
type User struct {
	gorm.Model
	Programs             []Program
	Email                string
	FirstName            string
	LastName             string
	Password             string
	PasswordConfirmation string
}

// BeforeSave - validations
func (u *User) BeforeSave() (err []error) {
	messages := make([]error, 0)

	// Presence Validations
	if u.FirstName == "" {
		messages = append(messages, errors.New("First Name cannot be blank"))
	}

	if u.Password == "" {
		messages = append(messages, errors.New("Password cannot be blank"))
	}

	if u.LastName == "" {
		messages = append(messages, errors.New("Last Name cannot be blank"))
	}

	if u.Email == "" {
		messages = append(messages, errors.New("Email cannot be blank"))
	}

	return messages
}
