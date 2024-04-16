package rest

import (
	"database/sql"
	"errors"
	"memorabilia/model"
	"memorabilia/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	userReq := &model.UserRegister{}
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	user, err := r.service.UserService.Register(userReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to register user", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "success", user)
}

func (r *Rest) Login(ctx *gin.Context) {
	userReq := &model.UserLogin{}
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}
	user, err := r.service.UserService.Login(userReq)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("username doesn't exists")
		}
		response.Error(ctx, http.StatusInternalServerError, "failed to login", err)
		return
	}
	response.Success(ctx, http.StatusOK, "success", user)
}

func (r *Rest) UpdateProfile(ctx *gin.Context) {
	var userReq model.UpdateUser

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	user, err := r.service.UserService.UpdateUser(ctx, &userReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update profile", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}

func (r *Rest) UploadProfilePicture(ctx *gin.Context) {
	profilePicture, err := ctx.FormFile("profilePicture")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	user, err := r.service.UserService.UploadProfilePicture(ctx, profilePicture)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to upload image", err)
		return
	}

	response.Success(ctx, http.StatusOK, "success", user)
}
