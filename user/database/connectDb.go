package database

import (
	"fmt"
	"log"

	models "github.com/A-junaid-K/pixel_vogue/user/models/request"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDb() {
	var err error
	dsn := "host=localhost user=ajk password=11012005 owner=pixelvogue port=5432 sslmode=disable"
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
