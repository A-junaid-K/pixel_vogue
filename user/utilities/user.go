package utilities

import (
	"github.com/A-junaid-K/pixel_vogue/user/database"
	models "github.com/A-junaid-K/pixel_vogue/user/models/request"
)

func FindUserById(Id int)(models.User,error){
	var user models.User

	if err := database.DB.Find(&user, Id).Error; err != nil{
		return models.User{},err
	}
	return user, nil
}
