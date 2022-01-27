package model

type Movie struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Comments    []Comment
}
