package getperson

import (
	"net/http"

	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/LOCNNIL/golang-rinha-api/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetPerson(repo *repository.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {
		person_request := GetPersonRequest{}

		if err := ctx.Bind(&person_request); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}

		// UUID validation
		if _, err := uuid.Parse(person_request.Id); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		person := models.People{
			Id: person_request.Id,
		}

		if tx := repo.FindPerson(&person); tx.Error != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Person Not Found.")
		}

		person_response := NewPersonResponse(person)

		return ctx.JSON(http.StatusOK, person_response)
	}
}
