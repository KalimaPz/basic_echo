package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Product struct {
	Uuid      string  `json:"uuid"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
	Catergory string  `json:"catergory"`
}

var productList []Product = []Product{
	{Uuid: "d62f4664-d60f-44b3-8c24-7a0fa0d09792", Name: "RTX 3070", Price: 17000, Catergory: "GPU"},
	{Uuid: "9da280f7-ecf9-49a2-8e3f-5b0330074965", Name: "RTX 3060", Price: 12000, Catergory: "GPU"},
	// {Uuid: "7ac84b3a-a011-42cc-96ac-e1ce146ce3a0", Name: "Ryzen 5 5600x", Price: 6190, Catergory: "CPU"},
	// {Uuid: "ecd3c4c0-a535-4afc-99c1-eb35544e8928", Name: "Ryzen 7 7600x", Price: 12000, Catergory: "CPU"},
	// {Uuid: "69f8e8c9-b3bd-43ec-8d07-abbcc206b3e4", Name: "Core i5 8400", Price: 8000, Catergory: "CPU"},
}

func GetProducts() (res []Product, err error) {
	res = productList
	return res, err
}

func CreateNewProduct(newProduct Product) (err error) {
	uuid := uuid.New()

	for _, product := range productList {

		if strings.EqualFold(product.Name, newProduct.Name) {
			fmt.Println("catch")
			err = errors.New("exception : this product already create")
			break

		} else {
			newProduct.Uuid = uuid.String()
			productList = append(productList, newProduct)
			err = nil
		}

	}
	return err

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
