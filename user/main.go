package main

import "github.com/go-playground/validator/v10"

type User struct {
	Username string `validate:"required,min=3,max=20"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"gte=0,lte=130"`
}

func main() {
	validator.New()
}
