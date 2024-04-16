package entity

import (
	"time"

	"github.com/google/uuid"
)

type Diary struct {
	ID        uuid.UUID       `json:"id"`
	UserID    uuid.UUID       `json:"userId"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	Pictures  []*DiaryPicture `json:"pictures"`
	CreatedAt time.Time       `json:"createdAt"`
}
