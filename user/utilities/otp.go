package utilities

import (
	"math/rand"
	"net/smtp"
	"os"

	"github.com/A-junaid-K/pixel_vogue/user/database"
	models "github.com/A-junaid-K/pixel_vogue/user/models/request"
	"github.com/gin-gonic/gin"
)

func GenerateOTP() int {
	return rand.Intn(9000) + 1000
}

func SendOtp(otp, email string) error {
	auth := smtp.PlainAuth("", os.Getenv("email"), os.Getenv("password"), "smtp.gmail.com")
	to := []string{email}
	message := "Subject: Otp verification\nyour verification otp is " + otp
	return smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("email"), to, []byte(message))
}

type OtpVerifiaction struct {
	Email string `json:"email"`
	Otp   int    `json:"otp"`
}

func OtpVerification(c *gin.Context) {
	// otp and geting from user
	var otp OtpVerifiaction
	if err := c.Bind(&otp); err != nil {
		resp := Response{
			StatusCode: 422,
			Error:      "failed to parse request body. Please ensure it's valid JSON",
			Data:       nil,
		}
		ResponseResult(c, resp)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", otp.Email).First(&user).Error; err != nil {
		resp := Response{
			StatusCode: 404,
			Error:      "User not found",
			Data:       nil,
		}
		ResponseResult(c, resp)
		return
	}

	// checking the user already validate or not.
	// if user.Validate {
	// 	resp := utilities.Response{
	// 		StatusCode: 409,
	// 		Err:        "User is already verified. Please log in directly",
	// 		Data:       nil,
	// 	}
	// 	utilities.ResponseResult(c, resp)
	// 	return
	// }

	// checking the otp correct or not.
	if otp.Otp != user.Otp {
		resp := Response{
			StatusCode: 400,
			Error:      "Invalid OTP entered. Please check your OTP and try again.",
			Data:       nil,
		}
		ResponseResult(c, resp)
		return
	}
	// if the otp is correct the value in database validate column is updating to true
	database.DB.Model(&models.User{}).Where("id = ?", user.Id).Update("validate", true)

	resp := Response{
		StatusCode: 200,
		Error:      nil,
		Data:       "Successfully created user",
	}
	ResponseResult(c, resp)
}
