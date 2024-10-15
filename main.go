package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mehmetolgundev/nba-project/domain/match"
	"github.com/mehmetolgundev/nba-project/infra"
)

func main() {
	e := echo.New()
	dbConnection := infra.NewConnection()
	matchRepository := match.NewRepository(dbConnection)
	matchService := match.NewService(matchRepository)

	e.GET("/api/matches", func(c echo.Context) error {
		currentTime, err := strconv.Atoi(c.QueryParam("currentTime"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		fixtures := matchService.GetMatches(context.Background(), currentTime)
		return c.JSON(http.StatusOK, fixtures)

	})

	e.File("/", "templates/index.html")
	e.Start(":8080")
}
