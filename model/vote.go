package model

type Vote struct {
	Id      int64   `json:"id"`
	User    User    `json:"user"`
	Rating  float64 `json:"rating"`
	MovieID int     `json:"movie_id"`
}
