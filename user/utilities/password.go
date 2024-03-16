package utilities

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Hashing password
func Hashpassword(password string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password")
	}
	return string(hashedPass)
}

// -----Compare Hashed password and User Entered password
func ComapreHashPassword(pass, hashedpass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass)); err != nil {
		return false
	} else {
		return true
	}
}
