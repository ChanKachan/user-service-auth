package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/database"
	"user-service/internal/logger"
	"user-service/internal/model"
	"user-service/repository"
)

func PostRegisterEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var employee model.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	rep := repository.NewEmployeeRepository(database.ConnectDataBase(), logger.InitLogger())
	str, err := rep.CreateEmployee(employee)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	json.NewEncoder(w).Encode(model.MessageResponse{str, http.StatusCreated})
}
