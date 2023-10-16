package countpeople

import (
	"net/http"
	"strconv"

	"github.com/LOCNNIL/golang-rinha-api/app/database/repository"
	"github.com/labstack/echo/v4"
)

func GetPerson(repo *repository.Repository) func(echo.Context) error {
	return func(ctx echo.Context) error {

		num, err := repo.GetPeopleCount()
		if err != nil {
			ctx.String(http.StatusOK, "0")
		}

		person_num := strconv.FormatInt(num, 10)
		return ctx.String(http.StatusOK, person_num)
	}
}
