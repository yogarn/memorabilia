package service

import (
	"memorabilia/internal/repository"
)

type Service struct {
	DiaryService IDiaryService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		DiaryService: NewDiaryService(repository.DiaryRepository),
	}
}
