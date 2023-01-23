package services

import (
	"basic_echo/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func GetProducts(ctx echo.Context) error {
	var data, err = repository.GetProducts()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
		})

	} else {
		return ctx.JSON(http.StatusOK, Response{
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
		return ctx.JSON(http.StatusInternalServerError, Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   nil,
		})

	} else {
		err := repository.CreateNewProduct(*reqBody)
		fmt.Println(err)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, Response{
				Status: false,
				Code:   http.StatusInternalServerError,
				Data:   err.Error(),
			})

		} else {

			return ctx.JSON(http.StatusOK, Response{
				Status: true,
				Code:   http.StatusOK,
				Data:   *reqBody,
			})

		}
	}

}

func GetProductByCatergory(ctx echo.Context) error {
	data, err := repository.GetProductByCatergory(ctx.QueryParam("catergory"))

	if err != nil {
		return ctx.JSON(http.StatusOK, Response{

			Status: false,
			Code:   http.StatusInternalServerError,
			Data:   err.Error(),
		})
	} else {
		return ctx.JSON(http.StatusOK, Response{

			Status: true,
			Code:   http.StatusOK,
			Data:   data,
		})
	}
}
