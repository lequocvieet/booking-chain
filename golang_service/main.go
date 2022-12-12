package main

import (
	"fmt"
	"log"
	"net/http"

	"golang_service/controller"
	db "golang_service/database"
	"golang_service/middlewares"
	"golang_service/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("-----------------")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	DB := db.Init()
	model := models.New(DB)
	u := controller.NewUserController(&model)
	r := controller.NewRoomController(&model)
	router := mux.NewRouter()
	router.HandleFunc("/auth/login", u.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", u.Register).Methods(http.MethodPost)
	router.HandleFunc("/create-list-room", middlewares.CheckJwt(r.CreateListRoom)).Methods(http.MethodPost)
	router.HandleFunc("/delete-list-room", middlewares.CheckJwt(r.DeleteListRoom)).Methods(http.MethodPost)
	router.HandleFunc("/create-room", middlewares.CheckJwt(r.CreateRoom)).Methods(http.MethodPost)
	router.HandleFunc("/update-room-state", middlewares.CheckJwt(r.UpdateRoomState)).Methods(http.MethodPost)
	router.HandleFunc("/book-room", middlewares.CheckJwt(r.BookRoom)).Methods(http.MethodPost)
	router.HandleFunc("/cancel-book-room", middlewares.CheckJwt(r.CancelBookRoom)).Methods(http.MethodPost)
	router.HandleFunc("/check-in", middlewares.CheckJwt(u.CheckIn)).Methods(http.MethodPost)
	router.HandleFunc("/check-out", middlewares.CheckJwt(u.CheckOut)).Methods(http.MethodPost)
	router.HandleFunc("/request-payment", middlewares.CheckJwt(u.RequestPayment)).Methods(http.MethodPost)
	router.HandleFunc("/transfer-roomNFT", middlewares.CheckJwt(u.TransferRoomNFT)).Methods(http.MethodPost)
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
