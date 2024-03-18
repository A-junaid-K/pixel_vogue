package main

import (
	"log"

	"github.com/A-junaid-K/pixel_vogue/user/database"
	"github.com/A-junaid-K/pixel_vogue/user/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env file : %v", err)
	}
	database.ConnectDb()
}

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":8000")
}
