package rest

import (
	"memorabilia/model"
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

	diary, err := r.service.DiaryService.CreateDiary(&diaryReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create diary", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "diary created", diary)
}

func (r *Rest) GetDiaryById(ctx *gin.Context) {
	id := ctx.Param("id")
	diary, err := r.service.DiaryService.GetDiaryById(id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to get diary", err)
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
	id := ctx.Param("id")
	var diaryReq model.UpdateDiary
	if err := ctx.BindJSON(&diaryReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind json", err)
		return
	}
	diary, err := r.service.DiaryService.UpdateDiary(id, &diaryReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to update diary", err)
		return
	}
	response.Success(ctx, http.StatusOK, "success", diary)
}
