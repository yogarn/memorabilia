package repository

import (
	"database/sql"
)

type Repository struct {
	DiaryRepository        IDiaryRepository
	UserRepository         IUserRepository
	DiaryPictureRepository IDiaryPictureRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		DiaryRepository:        NewDiaryRepository(db),
		UserRepository:         NewUserRepository(db),
		DiaryPictureRepository: NewDiaryPictureRepository(db),
	}
}
