package model

import "github.com/google/uuid"

type CreateDiary struct {
	Title  string    `json:"title" binding:"required"`
	Body   string    `json:"body" binding:"required"`
	UserID uuid.UUID `json:"userId" binding:"required"`
}

type UpdateDiary struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
