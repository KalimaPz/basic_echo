package router

import (
	"basic_echo/services"

	"github.com/labstack/echo/v4"
)

func InitialRoutes(e *echo.Echo) {
	e.GET("/products", services.GetProducts)
}
