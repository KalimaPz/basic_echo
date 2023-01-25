package models

type Product struct {
	Uuid       string  `json:"uuid"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	Price      float32 `json:"price"`
	Category   string  `json:"catagory"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}
