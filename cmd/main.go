package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-service/internal/handler"
	"user-service/internal/logger"
)

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetInfoUser).Methods("GET")
	r.HandleFunc("users/register", handler.PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
