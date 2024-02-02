package model

type Response struct {
	Cats []Cat `json:"cats"`
}

type Cat struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Owner *Owner `json:"owner"`
}
