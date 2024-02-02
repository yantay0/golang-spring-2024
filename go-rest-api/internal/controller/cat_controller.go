package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	catService "github.com/yantay0/golang-spring-2024/go-rest-api/internal/service"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check endpoint")
	w.WriteHeader(http.StatusOK)

	appName := "go-rest-api about cats for Tsis-1\nGolang application development,Spring 2024"
	author := "yantay0"

	fmt.Fprintf(w, "Server is up and running on 8080\n")
	fmt.Fprintf(w, "About app: %s\n", appName)
	fmt.Fprintf(w, "Author: %s\n", author)
}

func GetCats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	cats := catService.GetCats()
	jsonResponse, err := json.Marshal(cats)

	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func GetCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	cat := catService.GetCat(params["id"])
	jsonResponse, err := json.Marshal(cat)

	if err != nil {
		return
	}
	w.Write(jsonResponse)

}
