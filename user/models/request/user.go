package models

import (
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username,omitempty" validate:"min=3,max=20"`
	Email      string `json:"email,omitempty" validate:"email"`
	Password   string `json:"password,omitempty" validate:"min=6"`
	Name       string
	Is_blocked bool
	Is_admin   bool
	Creator    bool
	Created_at time.Time
}

// type UserSignUp struct{
// 	UID int
// 	Username string
// 	Password string
// }
