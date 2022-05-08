package models

import "time"

const (
	ArticleService       = "article"
	ArticleCacheTimeHour = 24
)

type Article struct {
	ID        uint64    `json:"ID" groups:"public"`
	Author    string    `json:"author" groups:"public"`
	Title     string    `json:"title" groups:"public"`
	Body      string    `json:"body" groups:"public"`
	CreatedAt time.Time `json:"created_at" groups:"public"`
}
