package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DiaryPicture struct {
	ID        uuid.UUID `json:"id"`
	DiaryID   uuid.UUID `json:"diaryId"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

type DiaryPictureNull struct {
	ID        sql.NullString `json:"id"`
	DiaryID   sql.NullString `json:"diaryId"`
	Link      sql.NullString `json:"link"`
	CreatedAt sql.NullTime   `json:"createdAt"`
}
