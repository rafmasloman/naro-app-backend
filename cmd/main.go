package main

import (
	"fmt"
	"log"
	"naro-app-be/internal/config"
	"naro-app-be/internal/delivery/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := config.NewDatabase()

	router := gin.Default()

	boostrap := config.Bootstrap(db)

	route := route.RouteConfig{
		App:            router,
		AuthController: &boostrap.AuthController,
		UserController: &boostrap.UserController,
	}

	route.SetupRoute()

	fmt.Println("server running on : 8001")

	route.App.Run(":8001")
}
