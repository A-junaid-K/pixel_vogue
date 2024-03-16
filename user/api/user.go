package api

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/A-junaid-K/pixel_vogue/user/database"
	models "github.com/A-junaid-K/pixel_vogue/user/models/request"
	"github.com/A-junaid-K/pixel_vogue/user/utilities"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidId         = errors.New("invalid Id provided")
	ErrFailedToParseBody = errors.New("failed to parse body. Please check your JSON request")

	validate = validator.New()
)

func SignUp(c *gin.Context) {
	var userReq models.User
	if err := c.Bind(&userReq); err != nil {
		resp := utilities.Response{
			StatusCode: 400,
			Message:    ErrFailedToParseBody.Error(),
			Data:       nil,
			Error:      err,
		}
		utilities.ResponseResult(c, resp)
		return
	}

	//	validating input by the validator package
	if err := validate.Struct(userReq); err != nil {
		resp := utilities.Response{
			StatusCode: 400,
			Data:       "Inpur is not valid",
			Error:      err,
		}
		utilities.ResponseResult(c, resp)
		return
	}

	// Checking Username already exist
	if user, err := utilities.FindUserById(userReq.Id); err == nil {
		resp := utilities.Response{
			StatusCode: http.StatusForbidden,
			Message:    "This Username already exist",
			Data:       user,
		}
		utilities.ResponseResult(c, resp)
		return
	}

	// Checking Email already exist
	var user models.User
	result := database.DB.Where("email=?", userReq.Email).First(&user)
	if result.Error != nil {
		panic("failed to query database")
	}
	if result.RowsAffected > 0 {
		resp := utilities.Response{
			StatusCode: http.StatusForbidden,
			Message:    "This Email already exist",
			Data:       user,
		}
		utilities.ResponseResult(c, resp)
		return
	}

	// Username validation
	// if err := utilities.UsernameValidator(userReq.Username); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Email validation
	// validate.Struct(userReq)

	// if err := utilities.EmailValidation(userReq.Email); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// Password validation
	length, number, upper, special := utilities.PasswordValidator(userReq.Password)
	switch {
	case !length:
		fmt.Println("password must be 6 or more")
	case !number:
		fmt.Println("must include numbers")
	case !upper:
		fmt.Println("at least one big letter")
	case !special:
		fmt.Println("at least one special charecter")
	}

	//-----Hash password
	pass := utilities.Hashpassword(userReq.Password)

	// ----Generate OTP

	database.DB.Create(models.User{
		Username: userReq.Username,
		Email: userReq.Email,
	})

}
