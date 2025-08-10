package model

import "time"

type Article struct {
	Id        string     `json:"id"`
	AuthorId  string     `json:"author_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	CreatedAt *time.Time `json:"created_at"`
}
