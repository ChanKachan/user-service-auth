package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/model"
)

func WriteError(w http.ResponseWriter, statusCode int, errMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResponse := model.ErrorMessage{
		Error:        http.StatusText(statusCode),
		Code:         statusCode,
		ErrorMessage: errMessage,
	}

	json.NewEncoder(w).Encode(errResponse)
}
