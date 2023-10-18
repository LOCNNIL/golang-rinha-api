package main

import (
	"github.com/LOCNNIL/golang-rinha-api/app/database"
	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/environment"
	"github.com/LOCNNIL/golang-rinha-api/app/routes"
	"github.com/LOCNNIL/golang-rinha-api/app/validation"
	"github.com/labstack/echo/v4"
)

func main() {
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
