package utilities

import (
	"errors"
	"regexp"
	"unicode"
)

const (
	emailRegex    = `^[\w\-\.]+@([\w\-]+\.)+[\w\-]{2,4}$`
	usernameRegex = "^[A-Za-z][A-Za-z0-9_]{7,29}$"
	nameRegex     = "/^[A-Za-z]{3,28}$/"
	numberRegex   = `^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$`
)

// ------Username validation
func UsernameValidator(username string) error {

	regx := regexp.MustCompile(usernameRegex)
	if regx.MatchString(username) {
		return nil
	} else {
		return errors.New("username is not valid")
	}

}

// ------Email validation
func EmailValidator(email string) error {
	regx := regexp.MustCompile(emailRegex)
	if !regx.MatchString(email) {
		return errors.New("email is not valid")
	}
	return nil
}

// -------Name validation
func NameValidator(name string) error {
	regx := regexp.MustCompile(nameRegex)
	if !regx.MatchString(name) {
		return errors.New("name is not valid")
	}
	return nil
}

// -----Password Validation
func PasswordValidator(s string) (sixOrMore, number, upper, special bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default: 
			//return false, false, false, false
		}
	}
	sixOrMore = letters >= 6 && letters <= 16
	return
}
