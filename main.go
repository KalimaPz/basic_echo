package main

import (
	"basic_echo/repository"
	"basic_echo/router"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	repository.InitializeMongoConnection(e)
	router.CreateRouter(e)
	e.Logger.Fatal(e.Start(":3070"))

}
