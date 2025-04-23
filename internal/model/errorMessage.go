package model

type ErrorMessage struct {
	Error        string `json:"error"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}
