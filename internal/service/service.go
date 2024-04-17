package service

import (
	"memorabilia/internal/repository"
	"memorabilia/pkg/bcrypt"
	"memorabilia/pkg/jwt"
	"memorabilia/pkg/supabase"
)

type Service struct {
	DiaryService        IDiaryService
	UserService         IUserService
	DiaryPictureService IDiaryPictureService
	PeopleService       IPeopleService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		DiaryService:        NewDiaryService(repository.DiaryRepository, jwt),
		UserService:         NewUserService(repository.UserRepository, bcrypt, jwt, supabase),
		DiaryPictureService: NewDiaryPictureService(repository.DiaryPictureRepository, supabase),
		PeopleService:       NewPeopleService(repository.PeopleRepository, jwt),
	}
}
