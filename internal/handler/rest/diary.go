package rest

import (
	"errors"
	"memorabilia/model"
	customErrors "memorabilia/pkg/errors"
	"memorabilia/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateDiary(ctx *gin.Context) {
	var diaryReq model.CreateDiary
	if err := ctx.ShouldBindJSON(&diaryReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	diary, err := r.service.DiaryService.CreateDiary(ctx, &diaryReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create diary", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "diary created", diary)
}

func (r *Rest) GetDiaryById(ctx *gin.Context) {
	id := ctx.Param("diaryId")
	diary, err := r.service.DiaryService.GetDiaryById(id)
	if err != nil {
		if errors.Is(err, customErrors.ErrRecordNotFound) {
			response.Error(ctx, http.StatusNotFound, "diary not found", err)
		} else {
			response.Error(ctx, http.StatusInternalServerError, "failed to get diary", err)
		}
		return
	}
	response.Success(ctx, http.StatusOK, "success", diary)
}

func (r *Rest) GetDiary(ctx *gin.Context) {
	diaries, err := r.service.DiaryService.GetDiary()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get diaries", err)
		return
	}
	response.Success(ctx, http.StatusOK, "success", diaries)
}

func (r *Rest) UpdateDiary(ctx *gin.Context) {
	id := ctx.Param("diaryId")
	var diaryReq model.UpdateDiary
	if err := ctx.BindJSON(&diaryReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind json", err)
		return
	}
	diary, err := r.service.DiaryService.UpdateDiary(id, &diaryReq)
	if err != nil {
		if errors.Is(err, customErrors.ErrNoRowUpdated) {
			response.Error(ctx, http.StatusNotFound, "diary not updated", err)
		} else {
			response.Error(ctx, http.StatusInternalServerError, "failed to get diary", err)
		}
		return
	}
	response.Success(ctx, http.StatusOK, "success", diary)
}

func (r *Rest) DeleteDiary(ctx *gin.Context) {
	id := ctx.Param("diaryId")
	err := r.service.DiaryService.DeleteDiary(id)
	if err != nil {
		if errors.Is(err, customErrors.ErrNoRowDeleted) {
			response.Error(ctx, http.StatusNotFound, "diary not deleted", err)
		} else {
			response.Error(ctx, http.StatusInternalServerError, "failed to get diary", err)
		}
		return
	}
	response.Success(ctx, http.StatusOK, "success", id)
}
