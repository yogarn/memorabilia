package service

import (
	"fmt"
	"memorabilia/entity"
	"memorabilia/internal/repository"
	"memorabilia/model"
	"memorabilia/pkg/bcrypt"
	"memorabilia/pkg/jwt"
	"memorabilia/pkg/supabase"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(userReq *model.UserRegister) (*entity.User, error)
	Login(userReq *model.UserLogin) (*model.UserLoginResponse, error)
	GetUserById(id uuid.UUID) (*entity.User, error)
	UpdateUser(ctx *gin.Context, user *model.UpdateUser) (*model.UpdateUser, error)
	UploadProfilePicture(ctx *gin.Context, profilePicture *multipart.FileHeader) (*model.UpdateUser, error)
}

type UserService struct {
	UserRepository       repository.IUserRepository
	Bcrypt               bcrypt.Interface
	JWT                  jwt.Interface
	Supabase             supabase.Interface
	ProfilePictureBucket string
}

func NewUserService(userRepository repository.IUserRepository, bcrypt bcrypt.Interface, jwt jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		UserRepository:       userRepository,
		Bcrypt:               bcrypt,
		JWT:                  jwt,
		Supabase:             supabase,
		ProfilePictureBucket: "profilePicture",
	}
}

func (userService *UserService) Register(userReq *model.UserRegister) (*entity.User, error) {
	hashPassword, err := userService.Bcrypt.GenerateFromPassword(userReq.Password)
	if err != nil {
		return nil, err
	}
	userEntity := &entity.User{
		ID:       uuid.New(),
		Username: userReq.Username,
		Password: hashPassword,
		Name:     userReq.Name,
		Email:    userReq.Email,
		RoleID:   2, //default user role
	}
	user, err := userService.UserRepository.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserService) Login(userReq *model.UserLogin) (*model.UserLoginResponse, error) {
	user, err := userService.UserRepository.LoginUser(userReq)
	if err != nil {
		return nil, err
	}

	err = userService.Bcrypt.CompareAndHashPassword(user.Password, userReq.Password)
	if err != nil {
		return nil, err
	}

	token, err := userService.JWT.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	response := &model.UserLoginResponse{
		Username: user.Username,
		Token:    token,
	}
	return response, nil
}

func (userService *UserService) GetUserById(id uuid.UUID) (*entity.User, error) {
	user, err := userService.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userService *UserService) UpdateUser(ctx *gin.Context, userReq *model.UpdateUser) (*model.UpdateUser, error) {
	userId, err := userService.JWT.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}

	if userReq.Password != "" {
		hashPassword, err := userService.Bcrypt.GenerateFromPassword(userReq.Password)
		if err != nil {
			return nil, err
		}
		userReq.Password = hashPassword
	}

	result, err := userService.UserRepository.UpdateUser(userId, userReq)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userService *UserService) UploadProfilePicture(ctx *gin.Context, profilePicture *multipart.FileHeader) (*model.UpdateUser, error) {
	userId, err := userService.JWT.GetLoginUser(ctx)
	if err != nil {
		return nil, err
	}

	userLogin, err := userService.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	if userLogin.ProfilePicture != "" {
		err = userService.Supabase.Delete(userService.ProfilePictureBucket, userLogin.ProfilePicture)
		if err != nil {
			return nil, err
		}
	}

	profilePicture.Filename = fmt.Sprintf("%v-%v", time.Now().Format("2006-01-02_150405"), profilePicture.Filename)
	link, err := userService.Supabase.Upload(userService.ProfilePictureBucket, profilePicture)
	if err != nil {
		return nil, err
	}
	userLogin.ProfilePicture = link

	updateUser := &model.UpdateUser{
		ProfilePicture: link,
	}

	result, err := userService.UserRepository.UpdateUser(userLogin.ID, updateUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
