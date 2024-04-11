package rest

import (
	"fmt"
	"memorabilia/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")
	diary := routerGroup.Group("/diary")
	diary.POST("", r.CreateDiary)
	diary.GET("/", r.GetDiary)
	diary.GET("/:id", r.GetDiaryById)
	diary.PATCH("/:id", r.UpdateDiary)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
