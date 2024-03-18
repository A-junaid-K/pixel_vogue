package models

import (
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username,omitempty" validate:"min=3,max=20"`
	Email      string `json:"email,omitempty" validate:"email"`
	Password   string `json:"password,omitempty" validate:"min=6"`
	Name       string `json:"name,omitempty" validate:"min=3,max=20"`
	Is_blocked bool   `json:"isblocked" gorm:"default=false"`
	Is_admin   bool   `json:"admin" gorm:"default=false"`
	Creator    bool   `json:"creator" gorm:"default=false"`
	Created_at time.Time
	Otp int
}

// type UserSignUp struct{
// 	UID int
// 	Username string
// 	Password string
// }
