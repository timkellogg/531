package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/timkellogg/531/server/config"
)

// User - model
// TODO - add hashing function | don't store plaintext password
type User struct {
	gorm.Model
	Programs []Program `json:"programs"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Weight   uint      `json:",string"`
	Password string    `json:"password"`
}

// ValidateUser - validations
func (u *User) ValidateUser() (err []error) {
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

	// Uniqueness Validations
	user := User{}

	record := config.Database.DB.First(&user, User{Email: u.Email}).Count
	if record != nil {
		messages = append(messages, errors.New("Email has to be unique"))
	}

	return messages
}

// encryptPassword - encrypts password using bcrypt hashing algo
func (u *User) encryptPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return string(bytes), err
}

// decryptPassword - checks if hash matches decrypted password
func (u *User) decryptPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
