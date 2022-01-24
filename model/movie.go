package model

type Movie struct {
	id          int64
	name        string
	description string
	rating      float64
	comments    Comment
}
