package putperson

import (
	"fmt"
	"net/http"

	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/models"
	"github.com/labstack/echo/v4"
)

func PutPerson(repo *repository.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		person_request := PostPersonRequest{}

		if err := ctx.Bind(&person_request); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		if err := ctx.Validate(&person_request); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		person := models.People{
			Id:        person_request.Id,
			Nickname:  person_request.Nickname,
			Name:      person_request.Name,
			Birthdate: person_request.Birthdate,
			Stack:     person_request.Stack,
		}

		if s := len(person.Nickname); s > 32 {
			msg := fmt.Sprintf("Not allowed Nickname size: %d (maximum 32)", s)
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}

		if s := len(person.Name); s > 100 {
			msg := fmt.Sprintf("Not allowd name size:  %d (maximum 100)", s)
			return echo.NewHTTPError(http.StatusBadRequest, msg)
		}

		// TODO: test if the person doesn't exists yet

		if err := repo.UpsertPerson(&person); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		ctx.Response().Header().Set("Location", "/pessoas/"+person.Id)

		return ctx.JSON(http.StatusCreated, person)
	}
}
