package models

import (
	"errors"

	"log"

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

// SaveUser - performs pw encryption then saves record
func (u *User) SaveUser() error {
	encryptedPassword, err := Encrypt(u.Password)
	if err != nil {
		log.Fatal(err)
	}

	user := User{Password: encryptedPassword, Username: u.Username, Email: u.Email, Weight: u.Weight}
	config.Database.DB.Create(&user)

	return err
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

	recordNotFound := config.Database.DB.First(&user, User{Email: u.Email}).RecordNotFound()

	if !recordNotFound {
		messages = append(messages, errors.New("Email has to be unique"))
	}

	return messages
}
