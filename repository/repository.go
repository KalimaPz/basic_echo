package repository

import (
	"errors"
	"strings"
)

type Response struct {
	Status bool        `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type Product struct {
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Catergory string  `json:"catergory"`
}

var productList []Product = []Product{
	{Name: "RTX 3070", Price: 17000, Catergory: "GPU"},
	{Name: "RTX 3060", Price: 12000, Catergory: "GPU"},
	{Name: "Ryzen 5 5600x", Price: 6190, Catergory: "CPU"},
	{Name: "Ryzen 7 7600x", Price: 12000, Catergory: "CPU"},
	{Name: "Core i5 8400", Price: 8000, Catergory: "CPU"},
}

func GetProducts() (res []Product, err error) {
	res = productList
	return res, err
}

func CreateNewProduct(newProduct Product) (status bool) {
	productList = append(productList, newProduct)
	return true
}

func GetProductByCatergory(catergory string) (res []Product, err error) {
	var result []Product = nil
	for _, product := range productList {
		if strings.EqualFold(product.Catergory, catergory) {
			result = append(result, product)
		}
	}

	if result == nil {
		err = errors.New("exception : cannot find catergory")
	}

	return result, err
}
