package service

import (
	"fmt"
	"memorabilia/entity"
	"memorabilia/internal/repository"
	"memorabilia/pkg/supabase"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type IDiaryPictureService interface {
	AddDiaryPicture(diaryID uuid.UUID, profilePicture *multipart.FileHeader) (*entity.DiaryPicture, error)
	DeleteDiaryPicture(ID uuid.UUID) error
}

type DiaryPictureService struct {
	DiaryPictureRepository repository.IDiaryPictureRepository
	Supabase               supabase.Interface
	DiaryPictureBucket     string
}

func NewDiaryPictureService(diaryPictureRepository repository.IDiaryPictureRepository, supabase supabase.Interface) IDiaryPictureService {
	return &DiaryPictureService{
		DiaryPictureRepository: diaryPictureRepository,
		Supabase:               supabase,
		DiaryPictureBucket:     "diaryPicture",
	}
}

func (diaryPictureService *DiaryPictureService) AddDiaryPicture(diaryID uuid.UUID, profilePicture *multipart.FileHeader) (*entity.DiaryPicture, error) {
	profilePicture.Filename = fmt.Sprintf("%v-%v", time.Now().Format("2006-01-02_150405"), profilePicture.Filename)
	link, err := diaryPictureService.Supabase.Upload(diaryPictureService.DiaryPictureBucket, profilePicture)
	if err != nil {
		return nil, err
	}

	diaryPicture := &entity.DiaryPicture{
		ID:      uuid.New(),
		DiaryID: diaryID,
		Link:    link,
	}

	diaryPicture, err = diaryPictureService.DiaryPictureRepository.AddDiaryPicture(diaryPicture)
	if err != nil {
		return diaryPicture, err
	}
	return diaryPicture, nil
}

func (diaryPictureService *DiaryPictureService) DeleteDiaryPicture(ID uuid.UUID) error {
	diaryPicture, err := diaryPictureService.DiaryPictureRepository.GetDiaryPictureById(ID)
	if err != nil {
		return err
	}

	err = diaryPictureService.Supabase.Delete(diaryPictureService.DiaryPictureBucket, diaryPicture.Link)
	if err != nil {
		return err
	}

	err = diaryPictureService.DiaryPictureRepository.DeleteDiaryPicture(ID)
	if err != nil {
		return err
	}
	return nil
}
