package controllers

import (
	"basic_echo/models"
	"basic_echo/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateNewProduct(ctx echo.Context) error {

	requestBody := new(models.Product)
	err := ctx.Bind(requestBody)

	if err == nil {

		res := repository.RepCreateNewProduct(*requestBody)

		if res == nil {
			return ctx.JSON(http.StatusOK, models.Response{
				Status:  true,
				Message: "Product has been create",
				Data:    *requestBody,
			})
		} else {
			return ctx.JSON(http.StatusOK, models.Response{
				Status:    false,
				ErrorData: res,
			})

		}

	} else {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

}

func GetProducts(ctx echo.Context) error {

	products := repository.RepGetAllProducts()
	return ctx.JSON(http.StatusOK, models.Response{
		Status: true,
		Data:   products,
	})
}

func GetProducyByPriceRange(ctx echo.Context) error {
	var err error
	query_min := ctx.QueryParam("min")
	query_max := ctx.QueryParam("max")
	min, _ := strconv.ParseFloat(query_min, 32)
	max, _ := strconv.ParseFloat(query_max, 32)

	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal Server Error")

	} else {

		result := repository.RepGetAllProductByPriceRange(float32(min), float32(max))

		if min >= max {
			return ctx.JSON(http.StatusBadRequest, models.Response{
				Status:  false,
				Message: "Bad Request",
				ErrorData: models.ErrorData{
					Code:         4002,
					ErrorTitle:   "Invalid price range",
					ErrorMessage: "Min must less than Max",
				},
			})
		} else {
			return ctx.JSON(http.StatusOK, models.Response{
				Status: true,
				Data:   result,
			})
		}

	}

}
