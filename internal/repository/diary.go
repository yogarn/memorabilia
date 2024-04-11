package repository

import (
	"errors"
	"fmt"
	"memorabilia/entity"
	"memorabilia/model"
	"strings"

	"database/sql"
)

type IDiaryRepository interface {
	CreateDiary(diary *entity.Diary) (*entity.Diary, error)
	GetDiaryById(id string) (*entity.Diary, error)
	GetDiary() ([]*entity.Diary, error)
	UpdateDiary(id string, diary *model.UpdateDiary) (*model.UpdateDiary, error)
}

type DiaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) IDiaryRepository {
	return &DiaryRepository{db}
}

func (diaryRepository *DiaryRepository) CreateDiary(diary *entity.Diary) (*entity.Diary, error) {
	stmt := `INSERT INTO diary (id, title, body, createdAt)
VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err := diaryRepository.db.Exec(stmt, diary.ID, diary.Title, diary.Body)
	if err != nil {
		return diary, err
	}

	return diary, nil
}

func (diaryRepository *DiaryRepository) GetDiaryById(id string) (*entity.Diary, error) {
	diary := &entity.Diary{}
	stmt := `SELECT * FROM diary WHERE id = ?`
	row := diaryRepository.db.QueryRow(stmt, id)
	err := row.Scan(&diary.ID, &diary.Title, &diary.Body, &diary.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("models: no matching record found")
		} else {
			return nil, err
		}
	}
	return diary, nil
}

func (diaryRepository *DiaryRepository) GetDiary() ([]*entity.Diary, error) {
	diaries := []*entity.Diary{}
	stmt := `SELECT * FROM diary ORDER BY createdAt DESC`
	rows, err := diaryRepository.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		diary := &entity.Diary{}
		err := rows.Scan(&diary.ID, &diary.Title, &diary.Body, &diary.CreatedAt)
		if err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}
	return diaries, nil
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

	stmt := fmt.Sprintf("UPDATE diary SET %s WHERE id = ?", strings.Join(column, ", "))
	result, err := diaryRepository.db.Exec(stmt, values...)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected <= 0 {
		return nil, errors.New("no row affected")
	}
	return diary, nil
}
