package main

import (
	"fmt"

	"github.com/LOCNNIL/golang-rinha-api/app/database"
	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/environment"
	"github.com/LOCNNIL/golang-rinha-api/app/routes"
	"github.com/LOCNNIL/golang-rinha-api/app/validation"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		msg := fmt.Sprintf("Error, coundn't not load .env vars: %s", err)
		panic(msg)
	}
	database_connection := database.CreateConnection()
	database.Migrate(database_connection)

	repo := repository.Repository{
		DatabaseConnection: database_connection,
	}

	server := echo.New()
	validation.CreateValidator(server)
	routes.CreateRoutes(server, &repo)

	port := environment.GetEnvOrDie("APPLICATION_PORT")

	server.Logger.Fatal(server.Start(":" + port))
}
