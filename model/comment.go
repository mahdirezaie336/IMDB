package model

import (
	"fmt"
	"time"
)

type Comment struct {
	Id          int64     `json:"id"`
	User        string    `json:"user"`
	Comment     string    `json:"comment"`
	Approved    bool      `json:"approved"`
	CreatedAt   time.Time `json:"created_at"`
	MovieID     int       `json:"movie_id"`
	CommentBody string    `json:"comment_body"`
}

func (c *Comment) String() string {
	return fmt.Sprint(c.Comment, ":", c.CreatedAt)
}
