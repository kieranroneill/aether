package main

import (
	"aether/internal/constants"
	"aether/internal/routes"
	"aether/internal/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()

	// create the root files directory
	err := utils.CreateDir(constants.RootFileDirectory)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // allow any origin, obviously a major security loophole, but this is just an experiment :)

	// /files
	e.GET(constants.FilesRoute, routes.NewGetFilesRoute())
	// /files/upload
	e.POST(fmt.Sprint(constants.FilesRoute, constants.UploadRoute), routes.NewPostFilesUploadRoute())
	// /versions
	e.GET(constants.VersionsRoute, routes.NewGetVersionsRoute())

	// start the server
	err = e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		e.Logger.Fatal(err)
	}
}
