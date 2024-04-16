package entity

import (
	"time"

	"github.com/google/uuid"
)

type DiaryPicture struct {
	ID        uuid.UUID `json:"id"`
	DiaryID   uuid.UUID `json:"diaryId"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}
