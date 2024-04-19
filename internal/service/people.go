package service

import (
	"errors"
	"memorabilia/entity"
	"memorabilia/internal/repository"
	"memorabilia/model"
	"memorabilia/pkg/jwt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IPeopleService interface {
	CreatePeople(ctx *gin.Context, peopleModel *model.CreatePeople) (*entity.People, error)
	UpdatePeople(id string, peopleReq *model.UpdatePeople) (*model.UpdatePeople, error)
	DeletePeople(id string) error
}

type PeopleService struct {
	PeopleRepository repository.IPeopleRepository
	jwt              jwt.Interface
}

func NewPeopleService(peopleRepository repository.IPeopleRepository, jwt jwt.Interface) IPeopleService {
	return &PeopleService{
		PeopleRepository: peopleRepository,
		jwt:              jwt,
	}
}

func (peopleService *PeopleService) CreatePeople(ctx *gin.Context, peopleModel *model.CreatePeople) (*entity.People, error) {
	userId, err := peopleService.jwt.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}

	people := &entity.People{
		ID:          uuid.New(),
		UserID:      userId,
		Name:        peopleModel.Name,
		Description: peopleModel.Description,
		Relation:    peopleModel.Relation,
		CreatedAt:   time.Now().UTC(),
	}

	people, err = peopleService.PeopleRepository.CreatePeople(people)
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (peopleService *PeopleService) UpdatePeople(id string, peopleReq *model.UpdatePeople) (*model.UpdatePeople, error) {
	if peopleReq.Name == "" && peopleReq.Description == "" && peopleReq.Relation == "" {
		return peopleReq, errors.New("provide at least one field to update")
	}
	people, err := peopleService.PeopleRepository.UpdatePeople(id, peopleReq)
	if err != nil {
		return peopleReq, err
	}
	return people, nil
}

func (peopleService *PeopleService) DeletePeople(id string) error {
	err := peopleService.PeopleRepository.DeletePeople(id)
	if err != nil {
		return err
	}
	return nil
}
