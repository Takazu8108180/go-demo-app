package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	CreateArticleRequestBody struct {
		Title    string `json: "title"`
		Contents string `json: "contents"`
		Username string `json: "username"`
	}

	CreateArticleResponseBody struct {
		ID        uuid.UUID `json: "article_id"`
		Title     string    `json: "title"`
		Contents  string    `json: "contents"`
		Username  string    `json: "username"`
		CreatedAt time.Time `json:"created_at"`
	}

	SendNiceResponseBody struct {
		ID        uuid.UUID `json: "article_id"`
		Title     string    `json: "title"`
		Contents  string    `json: "contents"`
		Username  string    `json: "username"`
		Nice      int       `json: "nice"`
		CreatedAt time.Time `json: "created_at"`
	}
)
