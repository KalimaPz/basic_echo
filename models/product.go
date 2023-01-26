package models

type Product struct {
	Uuid       string  `json:"uuid"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	Price      float32 `json:"price"`
	Category   string  `json:"catagory"`
	Created_at int64   `json:"created_at"`
	Updated_at int64   `json:"updated_at"`
}
