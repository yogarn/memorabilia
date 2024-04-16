package repository

import (
	"errors"
	"fmt"
	"memorabilia/entity"
	"memorabilia/model"
	"strings"

	"database/sql"

	"github.com/google/uuid"
)

type IDiaryRepository interface {
	CreateDiary(diary *entity.Diary) (*entity.Diary, error)
	GetDiaryById(id string) (*entity.Diary, error)
	GetDiary() ([]*entity.Diary, error)
	UpdateDiary(id string, diary *model.UpdateDiary) (*model.UpdateDiary, error)
	DeleteDiary(id string) error
}

type DiaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) IDiaryRepository {
	return &DiaryRepository{db}
}

func (diaryRepository *DiaryRepository) CreateDiary(diary *entity.Diary) (*entity.Diary, error) {
	stmt := `INSERT INTO diaries (id, userId, title, body, createdAt)
VALUES(?, ?, ?, ?, UTC_TIMESTAMP())`
	tx, err := diaryRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(stmt, diary.ID, diary.UserID, diary.Title, diary.Body)
	if err != nil {
		tx.Rollback()
		return diary, err
	}

	err = tx.Commit()
	return diary, err
}

func (diaryRepository *DiaryRepository) GetDiaryById(id string) (*entity.Diary, error) {
	diary := &entity.Diary{}

	stmt := `SELECT * FROM diaries 
	LEFT JOIN diary_pictures ON diaries.id = diary_pictures.diaryId
	WHERE diaries.id = ? 
	`

	tx, err := diaryRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(stmt, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("models: no matching record found")
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		diaryPictureNull := &entity.DiaryPictureNull{}

		err := rows.Scan(&diary.ID, &diary.UserID, &diary.Title, &diary.Body, &diary.CreatedAt,
			&diaryPictureNull.ID, &diaryPictureNull.DiaryID, &diaryPictureNull.Link, &diaryPictureNull.CreatedAt)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		diaryPicture, err := diaryPictureNull.ValidatePictureNullString()
		if err != nil {
			return nil, err
		}

		if diaryPicture != nil {
			diary.Pictures = append(diary.Pictures, diaryPicture)
		}
	}

	err = tx.Commit()
	return diary, err
}

func (diaryRepository *DiaryRepository) GetDiary() ([]*entity.Diary, error) {
	diariesMap := map[uuid.UUID]*entity.Diary{}
	stmt := `
	SELECT * FROM diaries 
	LEFT JOIN diary_pictures ON diaries.id = diary_pictures.diaryId
	`

	tx, err := diaryRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(stmt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		diaryRow := &entity.Diary{}
		diaryPictureNull := &entity.DiaryPictureNull{}

		err := rows.Scan(&diaryRow.ID, &diaryRow.UserID, &diaryRow.Title, &diaryRow.Body, &diaryRow.CreatedAt,
			&diaryPictureNull.ID, &diaryPictureNull.DiaryID, &diaryPictureNull.Link, &diaryPictureNull.CreatedAt)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		diary, ok := diariesMap[diaryRow.ID]
		if !ok {
			diariesMap[diaryRow.ID] = diaryRow
			diary = diariesMap[diaryRow.ID]
		}

		diaryPicture, err := diaryPictureNull.ValidatePictureNullString()
		if err != nil {
			return nil, err
		}

		if diaryPicture != nil {
			diary.Pictures = append(diary.Pictures, diaryPicture)
		}
	}

	diaries := []*entity.Diary{}
	for _, diary := range diariesMap {
		diaries = append(diaries, diary)
	}
	err = tx.Commit()
	return diaries, err
}

func (diaryRepository *DiaryRepository) UpdateDiary(id string, diary *model.UpdateDiary) (*model.UpdateDiary, error) {
	var column []string
	var values []interface{}

	if diary.Title != "" {
		column = append(column, "title = ?")
		values = append(values, diary.Title)
	}
	if diary.Body != "" {
		column = append(column, "body = ?")
		values = append(values, diary.Body)
	}

	values = append(values, id)

	stmt := fmt.Sprintf("UPDATE diaries SET %s WHERE id = ?", strings.Join(column, ", "))
	tx, err := diaryRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(stmt, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if rowsAffected <= 0 {
		tx.Rollback()
		return nil, errors.New("no row updated")
	}

	err = tx.Commit()
	return diary, err
}

func (diaryRepository *DiaryRepository) DeleteDiary(id string) error {
	stmt := `DELETE FROM diaries WHERE id = ?`
	tx, err := diaryRepository.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(stmt, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected <= 0 {
		tx.Rollback()
		return errors.New("no row deleted")
	}

	err = tx.Commit()
	return err
}
