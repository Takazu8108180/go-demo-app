package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `json: "commend_id`
	ArticleID uuid.UUID `json: "article_id`
	Message   string    `json: "message`
	CreatedAt time.Time `json: "created_at`
}

type Article struct {
	ID          uuid.UUID `json: "article_id`
	Title       string    `json: "title`
	Contents    string    `json: "contents`
	Username    string    `json: "username`
	NiceNum     int       `json: "nice`
	CommentList []Comment `json: "comments"`
	CreatedAt   time.Time `json:"created_at"`
}
