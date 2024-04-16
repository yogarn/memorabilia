package repository

import (
	"database/sql"
	"memorabilia/entity"

	"github.com/google/uuid"
)

type IDiaryPictureRepository interface {
	AddDiaryPicture(diaryPicture *entity.DiaryPicture) (*entity.DiaryPicture, error)
	DeleteDiaryPicture(ID uuid.UUID) error
	GetDiaryPictureById(ID uuid.UUID) (*entity.DiaryPicture, error)
}

type DiaryPictureRepository struct {
	db *sql.DB
}

func NewDiaryPictureRepository(db *sql.DB) IDiaryPictureRepository {
	return &DiaryPictureRepository{db}
}

func (diaryPictureRepository *DiaryPictureRepository) AddDiaryPicture(diaryPicture *entity.DiaryPicture) (*entity.DiaryPicture, error) {
	stmt := `INSERT INTO diary_pictures (id, diaryId, link, createdAt)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`
	tx, err := diaryPictureRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(stmt, diaryPicture.ID, diaryPicture.DiaryID, diaryPicture.Link)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return diaryPicture, err
}

func (diaryPictureRepository *DiaryPictureRepository) DeleteDiaryPicture(ID uuid.UUID) error {
	stmt := `DELETE FROM diary_pictures WHERE id = ?`
	tx, err := diaryPictureRepository.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(stmt, ID)
	if err != nil {
		return err
	}
	return nil
}

func (diaryPictureRepository *DiaryPictureRepository) GetDiaryPictureById(ID uuid.UUID) (*entity.DiaryPicture, error) {
	stmt := `SELECT * FROM diary_pictures WHERE id = ?`
	tx, err := diaryPictureRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(stmt, ID)
	diaryPicture := &entity.DiaryPicture{}
	err = row.Scan(&diaryPicture.ID, &diaryPicture.DiaryID, &diaryPicture.Link, &diaryPicture.CreatedAt)
	if err != nil {
		return nil, err
	}
	return diaryPicture, nil
}
