package repository

import (
	"log"

	"github.com/yantay0/golang-spring-2024/go-rest-api/internal/model"
)

func PrepareResponse() []model.Cat {
	var cats []model.Cat

	cats = append(cats, model.Cat{ID: "1", Name: "Tom", Owner: &model.Owner{FirstName: "John", LastName: "Smith"}, Facts: "A cat can jump 5 times as high as it is tall."})
	cats = append(cats, model.Cat{ID: "2", Name: "Jerry", Owner: &model.Owner{FirstName: "Anna", LastName: "Lembke"}, Facts: "A cat can sprint at about thirty-one miles per hour."})
	cats = append(cats, model.Cat{ID: "3", Name: "Bob", Owner: &model.Owner{FirstName: "James", LastName: "Smith"}, Facts: "A cat cannot see directly under its nose."})
	log.Print(cats)
	return cats

}
