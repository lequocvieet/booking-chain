package test

import (
	"fmt"
	"golang_service/controller"
	db "golang_service/database"
	"golang_service/models"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	DB := db.Init()
	model := models.New(DB)
	r := controller.NewRoomController(&model)
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	router.Handle("/create-room", http.HandlerFunc(r.CreateRoom))
	//Token from login
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzA5NTgxMDksInVzZXJuYW1lIjoibWF5In0.OXkwkz77fiwBABwB8_n_p_zMgCFp6o-WyJih0imvXoA"
	body := "listRoomId=1&pricePerDay=100"
	req, err := http.NewRequest("POST", "/create-room", strings.NewReader(body))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestUpdateRoomState(t *testing.T) {
	DB := db.Init()
	model := models.New(DB)
	r := controller.NewRoomController(&model)
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	router.Handle("/update-room-state", http.HandlerFunc(r.UpdateRoomState))
	//Token from login
	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzA5NTgxMDksInVzZXJuYW1lIjoibWF5In0.OXkwkz77fiwBABwB8_n_p_zMgCFp6o-WyJih0imvXoA"
	body := "stateRoomId=4&listRoomId=1&roomId=4&startUpdate=2022-12-14&endUpdate=2022-12-16&state=Fix the air condition"
	req, err := http.NewRequest("POST", "/update-room-state", strings.NewReader(body))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}
