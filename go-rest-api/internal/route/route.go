package route

import (
	"log"

	"github.com/gorilla/mux"
	. "github.com/yantay0/golang-spring-2024/go-rest-api/internal/controller"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	log.Println("creating routes...")
	router.HandleFunc("/api/v1/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/api/v1/cats", GetCats).Methods("GET")
	router.HandleFunc("/api/v1/cats/{id}", GetCat).Methods("GET")

	return router
}
