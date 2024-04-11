package entity

import (
	"time"

	"github.com/google/uuid"
)

type Diary struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}
