package main

import (
	"authentication/database"
	"authentication/repository"
	service "authentication/services"
	"log"

	"github.com/labstack/echo/v4"
)

// @title Swagger Authentication API
// @version 1.0
// @description This is a authentication API.

// @contact.name Paulo Ferreira
// @contact.email jr@live.at

// @host localhost:8888
// @BasePath /
func main() {
	log.Println("Starting api server...")

	/*
		viper.SetConfigName("config")
		viper.AddConfigPath("../../config/")
		viper.SetConfigType("yml")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		} */

	db := database.NewPostgresDatabase()

	e := echo.New()
	SetRoutes(e, service.NewUserService(repository.NewUserRepository(db)))
	e.Logger.Fatal(e.Start(":80"))
}
