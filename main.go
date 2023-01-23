package main

import (
	"basic_echo/router"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	router.InitialRoutes(e)
	e.Logger.Fatal(e.Start(":3070"))

	summation(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func summation(numbers ...int) int {

	result := 0

	for _, v := range numbers {
		result += v
	}

	return result
}
