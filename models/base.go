package models

import "golang.org/x/crypto/bcrypt"

// Base - base model
type Base struct{}

// Encrypt string - encrypts string using bcrypt hashing algo
func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(bytes), err
}

// Decrypt - checks if hash matches decrypted string
func Decrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
