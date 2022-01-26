package model

type Movie struct {
	id          int64
	Name        string `json:"name"`
	Description string `json:"description"`
	rating      float64
	comments    []Comment
}
