package test

import (
	"fmt"
	"golang_service/controller"
	db "golang_service/database"
	"golang_service/models"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var err = godotenv.Load("local.env")
var router = mux.NewRouter()

func TestRegister(t *testing.T) {
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	DB := db.Init()
	model := models.New(DB)
	u := controller.NewUserController(&model)
	router.Handle("/auth/register", http.HandlerFunc(u.Register))
	body := "username=mm&email=bb@gmail.com&password=mm"
	req, err := http.NewRequest("POST", "/auth/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestLogin(t *testing.T) {
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	DB := db.Init()
	model := models.New(DB)
	u := controller.NewUserController(&model)
	router.Handle("/auth/login", http.HandlerFunc(u.Login))
	body := "username=vieet&password=vieet"
	req, err := http.NewRequest("POST", "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Errorf("Wrong request :%v", err)
	}
	response := executeRequest(req, router)

	fmt.Printf("res: %+v\n", response)

	checkResponseCode(t, http.StatusOK, response.Code)

}

func executeRequest(req *http.Request, router *mux.Router) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
