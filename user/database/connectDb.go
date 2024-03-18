package database

import (
	"fmt"
	"log"
	"os"

	models "github.com/A-junaid-K/pixel_vogue/user/models/request"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn:=os.Getenv("dsn")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	if err := DB.AutoMigrate(
		&models.User{},
	); err != nil {
		fmt.Printf("failed to auto migrate the model : %v",err)
		return
	}
}
