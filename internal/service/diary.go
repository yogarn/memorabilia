package service

import (
	"memorabilia/entity"
	"memorabilia/internal/repository"
	"memorabilia/model"

	"github.com/google/uuid"
)

type IDiaryService interface {
	CreateDiary(diaryReq *model.CreateDiary) (*entity.Diary, error)
	GetDiaryById(id string) (*entity.Diary, error)
	GetDiary() ([]*entity.Diary, error)
	UpdateDiary(id string, diary *model.UpdateDiary) (*model.UpdateDiary, error)
}

type DiaryService struct {
	DiaryRepository repository.IDiaryRepository
}

func NewDiaryService(diaryRepository repository.IDiaryRepository) IDiaryService {
	return &DiaryService{diaryRepository}
}

func (diaryService *DiaryService) CreateDiary(diaryReq *model.CreateDiary) (*entity.Diary, error) {
	diary := &entity.Diary{
		ID:    uuid.New(),
		Title: diaryReq.Title,
		Body:  diaryReq.Body,
	}
	diary, err := diaryService.DiaryRepository.CreateDiary(diary)
	if err != nil {
		return nil, err
	}
	return diary, nil
}

func (diaryService *DiaryService) GetDiaryById(id string) (*entity.Diary, error) {
	diary, err := diaryService.DiaryRepository.GetDiaryById(id)
	if err != nil {
		return nil, err
	}
	return diary, nil
}

func (diaryService *DiaryService) GetDiary() ([]*entity.Diary, error) {
	diaries, err := diaryService.DiaryRepository.GetDiary()
	if err != nil {
		return nil, err
	}
	return diaries, nil
}

func (diaryService *DiaryService) UpdateDiary(id string, diaryReq *model.UpdateDiary) (*model.UpdateDiary, error) {
	diary, err := diaryService.DiaryRepository.UpdateDiary(id, diaryReq)
	if err != nil {
		return nil, err
	}
	return diary, nil
}
