package repository

import (
	"log"

	"github.com/yantay0/golang-spring-2024/go-rest-api/internal/model"
)

func PrepareResponse() []model.Cat {
	var cats []model.Cat

	cats = append(cats, model.Cat{ID: "1", Name: "Tom", Owner: &model.Owner{FirstName: "John", LastName: "Smith"}})
	cats = append(cats, model.Cat{ID: "2", Name: "Jerry", Owner: &model.Owner{FirstName: "Anna", LastName: "Lembke"}})
	log.Print(cats)
	return cats

}
