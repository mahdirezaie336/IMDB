package model

import (
	"fmt"
	"time"
)

type Comment struct {
	id        int64
	user      User
	comment   string
	approved  bool
	createdAt time.Time
	movieID   int
}

func (c *Comment) String() string {
	return fmt.Sprint(c.comment, ":", c.createdAt)
}
