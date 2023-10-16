package routes

import (
	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/routes/people"
	"github.com/labstack/echo/v4"
)

func CreateRoutes(server *echo.Echo, repo *repository.Repository) {
	people.CreateRoutes(server, repo)
}
