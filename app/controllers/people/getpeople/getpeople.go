package getpeople

import (
	"net/http"

	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/models"
	"github.com/labstack/echo/v4"
)

func FindPeople(repo *repository.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		search_term := ctx.QueryParam("t")

		if search_term == "" {
			return ctx.JSON(http.StatusBadRequest, "Empty search term")
		}

		people := []models.People{}
		repo.FindPeople(&people, search_term)

		people_response := make([]FindPeopleResponse, len(people))

		for i, person := range people {
			people_response[i] = FindPeopleResponse{
				Id:        person.Id,
				Nickname:  person.Nickname,
				Name:      person.Name,
				Birthdate: person.Birthdate,
				Stack:     person.Stack,
			}
		}

		return ctx.JSON(http.StatusOK, people_response)
	}
}
