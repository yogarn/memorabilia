package repository

import (
	"database/sql"
)

type Repository struct {
	DiaryRepository IDiaryRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DiaryRepository: NewDiaryRepository(db),
	}
}
