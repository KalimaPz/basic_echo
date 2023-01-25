package repository

import (
	"basic_echo/models"
	"strings"

	"github.com/google/uuid"
)

var products []models.Product = []models.Product{
	{Uuid: "d62f4664-d60f-44b3-8c24-7a0fa0d09792", Name: "RTX 3071", Price: 17000, Category: "GPU", Brand: "MSI", Created_at: "25-01-2023"},
	{Uuid: "d62f4664-d60f-44b3-8c24-7a0fa0d09792", Name: "RTX 3072", Price: 12000, Category: "GPU", Brand: "MSI", Created_at: "25-01-2023"},
	{Uuid: "d62f4664-d60f-44b3-8c24-7a0fa0d09792", Name: "RTX 3073", Price: 15000, Category: "GPU", Brand: "MSI", Created_at: "25-01-2023"},
	{Uuid: "d62f4664-d60f-44b3-8c24-7a0fa0d09792", Name: "RTX 3074", Price: 19000, Category: "GPU", Brand: "MSI", Created_at: "25-01-2023"},
}

func RepCreateNewProduct(newProduct models.Product) interface{} {
	var err interface{}
	for _, p := range products {
		if strings.EqualFold(p.Name, newProduct.Name) {
			err = models.ErrorData{
				Code:         4001,
				ErrorTitle:   "Duplicate",
				ErrorMessage: "This item already has in store",
			}

		}
	}

	uuid := uuid.New()
	newProduct.Uuid = uuid.String()
	products = append(products, newProduct)

	return err
}

func RepGetAllProducts() (result []models.Product) {
	return products
}

func RepGetAllProductByPriceRange(min float32, max float32) []models.Product {

	var result []models.Product

	for _, p := range products {
		if p.Price >= min && p.Price <= max {
			result = append(result, p)
		}

	}

	return result
}

func RepGetProductByCategory() {}

func RepDeleteProduct() {}

func RepGetSortedProductsByPrice() {

}
