package models

type ErrorData struct {
	Code         int16  `json:"code"`
	ErrorTitle   string `json:"error_title"`
	ErrorMessage string `json:"error_msg"`
}
