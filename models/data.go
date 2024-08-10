package models

import (
	"log"
	"time"

	"github.com/google/uuid"
)

func ParseUUID(id string) uuid.UUID {
	uuid, err := uuid.Parse(id)

	if err != nil {
		log.Fatal(err)
	}

	return uuid
}

var (
	Comment1 = Comment{
		CommentID: ParseUUID("0380cb6b-97fd-8965-a167-05b3c5de0d57"),
		ArticleID: ParseUUID("1ef53bc7-9c08-060f-7889-48462a352a34"),
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	Comment2 = Comment{
		CommentID: ParseUUID("3389ae0c-8dda-f9c2-bd88-0dde1f361e96"),
		ArticleID: ParseUUID("1ef53bc7-9c08-060f-7889-48462a352a34"),
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
)
var (
	Article1 = Article{
		ID:          ParseUUID("1ef53bc7-9c08-060f-7889-48462a352a34"),
		Title:       "first article",
		Contents:    "This is the test article.",
		Username:    "saki",
		NiceNum:     1,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}
	Article2 = Article{
		ID:        ParseUUID("1a6fd54f-c4c6-c8a6-995b-833c01c3064a"),
		Title:     "second article",
		Contents:  "This is the test article.",
		Username:  "saki",
		NiceNum:   2,
		CreatedAt: time.Now(),
	}
)
