package main

import (
	"aether/internal/constants"
	"aether/internal/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

var Version string

func main() {
	e := echo.New()

	e.GET("/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.Application{
			Environment: os.Getenv("ENVIRONMENT"),
			Name:        constants.Name,
			Version:     Version,
		})
	})

	// start the server
	err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
