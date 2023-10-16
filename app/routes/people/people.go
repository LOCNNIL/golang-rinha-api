package people

import (
	"github.com/LOCNNIL/golang-rinha-api/app/controllers/people/countpeople"
	"github.com/LOCNNIL/golang-rinha-api/app/controllers/people/getpeople"
	"github.com/LOCNNIL/golang-rinha-api/app/controllers/people/getperson"
	"github.com/LOCNNIL/golang-rinha-api/app/controllers/people/putperson"
	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/labstack/echo/v4"
)

func CreateRoutes(server *echo.Echo, repo *repository.Repository) {
	server.POST("/pessoas", putperson.PutPerson(repo))
	server.GET("/pessoas/:id", getperson.GetPerson(repo))
	server.GET("/pessoas", getpeople.FindPeople(repo))
	server.GET("/contagem-pessoas", countpeople.GetPerson(repo))
}
