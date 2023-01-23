package repository

import (
	"errors"
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
}

func GetProducts() (res []Product, err error) {
	err = errors.New("Exception : New Exception")
	return res, err
}
