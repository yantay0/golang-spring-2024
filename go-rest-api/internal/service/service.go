package service

import (
	"github.com/yantay0/golang-spring-2024/go-rest-api/internal/model"
	catRepository "github.com/yantay0/golang-spring-2024/go-rest-api/internal/repository"
)

func GetCats() []model.Cat {
	cats := catRepository.PrepareResponse()
	return cats
}

func GetCat(id string) model.Cat {
	cats := catRepository.PrepareResponse()
	var cat model.Cat
	for _, item := range cats {
		if item.ID == id {
			cat = item
		}
	}
	return cat
}
