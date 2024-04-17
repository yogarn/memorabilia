package rest

import (
	"memorabilia/model"
	"memorabilia/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreatePeople(ctx *gin.Context) {
	peopleReq := &model.CreatePeople{}
	if err := ctx.ShouldBindJSON(peopleReq); err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request", err)
		return
	}

	people, err := r.service.PeopleService.CreatePeople(ctx, peopleReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "failed to create people", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "success", people)
}
