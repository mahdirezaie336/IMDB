package model

type Vote struct {
	id      int64
	user    User
	rating  float64
	movieID int
}
