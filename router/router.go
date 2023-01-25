package router

import (
	"basic_echo/controllers"

	"github.com/labstack/echo/v4"
)

func CreateRouter(e *echo.Echo) {
	version := "/v1"
	productEndpoint := e.Group("/api" + version)
	productEndpoint.GET("/product", controllers.GetProducts)
	productEndpoint.POST("/product", controllers.CreateNewProduct)
	productEndpoint.GET("/find_with_price_range", controllers.GetProducyByPriceRange)

}
