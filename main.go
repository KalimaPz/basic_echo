package main

import (
	"basic_echo/router"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	router.CreateRouter(e)
	e.Logger.Fatal(e.Start(":3070"))

}
