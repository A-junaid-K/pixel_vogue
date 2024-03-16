package utilities

import "crypto/rand"

func GenerateOTP()string{
	rand.Int(9000)
	return "otp"
}