package services

import (
	"basic_echo/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SayHello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World 4")

}

func GetProducts(ctx echo.Context) error {
	var data, err = repository.GetProducts()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, repository.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
		})

	} else {
		return ctx.JSON(http.StatusOK, repository.Response{
			Status: true,
			Code:   http.StatusOK,
			Data:   data,
		})
	}

}

func CreateNewProduct(ctx echo.Context) error {

	reqBody := new(repository.Product)
	err := ctx.Bind(reqBody)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, repository.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   nil,
		})

	} else {
		status := repository.CreateNewProduct(*reqBody)

		if status {
			return ctx.JSON(http.StatusOK, repository.Response{
				Status: true,
				Code:   http.StatusOK,
				Data:   *reqBody,
			})
		} else {
			return ctx.JSON(http.StatusInternalServerError, repository.Response{
				Status: false,
				Code:   http.StatusInternalServerError,
				Data:   nil,
			})

		}
	}

}

func GetProductByCatergory(ctx echo.Context) error {
	data, err := repository.GetProductByCatergory(ctx.QueryParam("catergory"))

	if err != nil {
		return ctx.JSON(http.StatusOK, repository.Response{

			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
		})
	} else {
		return ctx.JSON(http.StatusOK, repository.Response{

			Status: true,
			Code:   http.StatusOK,
			Data:   data,
		})
	}
}
