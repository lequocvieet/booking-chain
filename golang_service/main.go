package main

import (
	"fmt"
	"log"
	"net/http"

	db "golang_service/database"
	"golang_service/handler"
	"golang_service/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	DB := db.Init()
	h := handler.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", h.Register).Methods(http.MethodPost)
	router.HandleFunc("/create-list-room", middlewares.CheckJwt(h.CreateListRoom)).Methods(http.MethodPost)
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
