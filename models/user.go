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
	Username             string
	Email                string
	Weight               uint
	Password             string
	PasswordConfirmation string
}

// BeforeSave - validations
func (u *User) BeforeSave() (err []error) {
	messages := make([]error, 0)

	// Presence Validations
	if u.Password == "" {
		messages = append(messages, errors.New("Password cannot be blank"))
	}

	if u.Email == "" {
		messages = append(messages, errors.New("Email cannot be blank"))
	}

	if u.Username == "" {
		messages = append(messages, errors.New("Username cannot be blank"))
	}

	return messages
}
