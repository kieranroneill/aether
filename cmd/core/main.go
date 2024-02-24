package main

import (
	"aether/internal/constants"
	"aether/internal/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

var Version string

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
	}))

	// /files/upload
	e.POST(fmt.Sprint(constants.FilesRoute, constants.UploadRoute), routes.NewFilesUploadRoute())
	// /versions
	e.GET(constants.VersionsRoute, routes.NewVersionsRoute(Version))

	// start the server
	err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
