package main

import (
	"net/http"

	"github.com/yantay0/golang-spring-2024/go-rest-api/internal/route"
)

func main() {
	router := route.SetupRoutes()
	http.ListenAndServe(":8080", router)
}
