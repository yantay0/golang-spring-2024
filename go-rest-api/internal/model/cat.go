package model

type Response struct {
	Cats []Cat `json:"cats"`
}

type Cat struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner *Owner `json:"owner"`
	Facts string `json:"facts"`
}
