package model

type User struct {
	Id       int64  `json:"id"`
	Role     int64  `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}
