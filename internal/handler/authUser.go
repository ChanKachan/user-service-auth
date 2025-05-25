package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"user-service/internal/database"
	"user-service/internal/logger"
	"user-service/internal/model"
	"user-service/repository"
)

// GetInfoUser возращает при успешном вводе JSON файл при успешном поиске
func GetInfoUser(w http.ResponseWriter, r *http.Request) {
	logger := logger.InitLogger()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	rep := repository.NewUserRepository(database.ConnectDataBase(), logger)
	user, err := rep.GetUserByID(id)
	if err != nil {
		logger.Error("function GetUserByID", zap.Error(err))
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&model.User{
		ID: id, Name: user.Name, Patronymic: user.Patronymic, Surname: user.Surname, Email: user.Email, PhoneNumber: user.PhoneNumber, Login: user.Login,
	})
}

// Добавляем нового юзера, при введение корректных данных
func PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user *model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	rep := repository.NewUserRepository(database.ConnectDataBase(), logger.InitLogger())
	user, err = rep.CreateUser(user)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginRequest *model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	rep := repository.NewLoginRepository(database.ConnectDataBase(), logger.InitLogger())
	loginResponse, err := rep.Authenticate(loginRequest)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&loginResponse)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	rep := repository.NewUserRepository(database.ConnectDataBase(), logger.InitLogger())
	err = rep.UpdateUser(user)

	if err != nil {
		WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	json.NewEncoder(w).Encode(&model.User{
		ID: user.ID, Name: user.Name, Patronymic: user.Patronymic, Surname: user.Surname, Email: user.Email, PhoneNumber: user.PhoneNumber, Login: user.Login,
	})
}
