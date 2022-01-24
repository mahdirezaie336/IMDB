package model

import "time"

type Comment struct {
	id        int64
	user      User
	comment   string
	approved  bool
	createdAt time.Time
	movieID   int
}
