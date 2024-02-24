package main

import (
	"aether/internal/constants"
	"aether/internal/files"
	"aether/internal/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()

	// create the root files directory
	writeError := files.CreateRootFilesDirectory()
	if writeError != nil {
		e.Logger.Fatal(writeError.Error)
	}

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // allow any origin, obviously a major security loophole, but this is just an experiment :)

	// /files/upload
	e.POST(fmt.Sprint(constants.FilesRoute, constants.UploadRoute), routes.NewFilesUploadRoute())
	// /versions
	e.GET(constants.VersionsRoute, routes.NewVersionsRoute())

	// start the server
	err := e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
