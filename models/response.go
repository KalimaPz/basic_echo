package models

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"msg"`

	ErrorData interface{} `json:"error_data"`
	Data      interface{} `json:"data"`
}
