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

func (diaryPictureNull DiaryPictureNull) ValidatePictureNullString() (*DiaryPicture, error) {
	if diaryPictureNull.ID.Valid {
		diaryPicture := &DiaryPicture{}
		if diaryPictureNull.ID.Valid {
			parsedUUID, err := uuid.Parse(diaryPictureNull.ID.String)
			if err != nil {
				return nil, err
			}

			diaryPicture.ID = parsedUUID
		}

		if diaryPictureNull.DiaryID.Valid {
			parsedUUID, err := uuid.Parse(diaryPictureNull.DiaryID.String)
			if err != nil {
				return nil, err
			}

			diaryPicture.DiaryID = parsedUUID
		}

		if diaryPictureNull.Link.Valid {
			diaryPicture.Link = diaryPictureNull.Link.String
		}

		if diaryPictureNull.CreatedAt.Valid {
			diaryPicture.CreatedAt = diaryPictureNull.CreatedAt.Time
		}
		return diaryPicture, nil
	}
	return nil, nil
}
