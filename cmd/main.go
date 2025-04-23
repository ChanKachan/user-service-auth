package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"user-service/internal/auth"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/logger"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	config.Init(jwtKey)
	logger.Info("JTW key created successfully")

	logger.Info("Launch user handlers")
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetInfoUser).Methods("GET")
	r.HandleFunc("/users/register", handler.PostRegisterUser).Methods("POST")
	r.HandleFunc("/users/auth", handler.PostLogin).Methods("POST")
	r.HandleFunc("/users/updateUser", auth.JWTAuthMiddleware(handler.PutUser)).Methods("PUT")

	logger.Info("User handlers work correctly")

	r.HandleFunc("/users/employee", auth.JWTAuthMiddleware(handler.PostRegisterEmployee)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
