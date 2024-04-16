package service

import (
	"memorabilia/entity"
	"memorabilia/internal/repository"
	"memorabilia/model"
	"memorabilia/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IDiaryService interface {
	CreateDiary(ctx *gin.Context, diaryReq *model.CreateDiary) (*entity.Diary, error)
	GetDiaryById(id string) (*entity.Diary, error)
	GetDiary() ([]*entity.Diary, error)
	UpdateDiary(id string, diary *model.UpdateDiary) (*model.UpdateDiary, error)
	DeleteDiary(id string) error
}

type DiaryService struct {
	DiaryRepository repository.IDiaryRepository
	jwt             jwt.Interface
}

func NewDiaryService(diaryRepository repository.IDiaryRepository, jwt jwt.Interface) IDiaryService {
	return &DiaryService{
		DiaryRepository: diaryRepository,
		jwt:             jwt,
	}
}

func (diaryService *DiaryService) CreateDiary(ctx *gin.Context, diaryReq *model.CreateDiary) (*entity.Diary, error) {
	userId, err := diaryService.jwt.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}
	diary := &entity.Diary{
		ID:     uuid.New(),
		UserID: userId,
		Title:  diaryReq.Title,
		Body:   diaryReq.Body,
	}
	diary, err = diaryService.DiaryRepository.CreateDiary(diary)
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

func (diaryService *DiaryService) DeleteDiary(id string) error {
	err := diaryService.DiaryRepository.DeleteDiary(id)
	if err != nil {
		return err
	}
	return nil
}
