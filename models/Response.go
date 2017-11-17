package models

// Response json model
type Response struct {
	Code   int         `json:"code"`
	Error  string      `json:"error"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Data   interface{} `json:"data"`
}
