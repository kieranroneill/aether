package routes

import (
	"aether/internal/types"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func NewVersionsRoute() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, types.VersionsResponse{
			Environment: os.Getenv("ENVIRONMENT"),
			Name:        os.Getenv("NAME"),
			Version:     os.Getenv("VERSION"),
		})
	}
}
