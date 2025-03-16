package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"user-service/internal/config"
	"user-service/internal/handler"
	"user-service/internal/logger"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	config.Init(&jwtKey)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetInfoUser).Methods("GET")
	r.HandleFunc("/users/register", handler.PostUser).Methods("POST")
	r.HandleFunc("/users/auth", handler.PostLogin).Methods("POST")
	r.HandleFunc("/users/updateUser", handler.PutUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
