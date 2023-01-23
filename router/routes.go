package router

import (
	"basic_echo/services"

	"github.com/labstack/echo/v4"
)

func InitialRoutes(e *echo.Echo) {

	productsGroup := e.Group("/products")
	productsGroup.GET("", services.GetProducts)
	productsGroup.GET("/getProductByCatergory", services.GetProductByCatergory)
	productsGroup.POST("", services.CreateNewProduct)

}
