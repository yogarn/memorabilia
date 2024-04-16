package rest

import (
	"memorabilia/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) AddDiaryPicture(ctx *gin.Context) {
	id := ctx.Param("diaryId")
	diaryPictureFile, err := ctx.FormFile("diaryPicture")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to bind form data", err)
		return
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "failed to parse uuid", err)
		return
	}

	diaryPicture, err := r.service.DiaryPictureService.AddDiaryPicture(parsedId, diaryPictureFile)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to add picture to diary", err)
	}
	response.Success(ctx, http.StatusOK, "success", diaryPicture)
}
