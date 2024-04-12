package rest

import (
	"fmt"
	"memorabilia/internal/service"
	"memorabilia/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func MountDiary(routerGroup *gin.RouterGroup, r *Rest) {
	diary := routerGroup.Group("/diary")
	diary.POST("", r.CreateDiary)
	diary.GET("/", r.GetDiary)
	diary.GET("/:id", r.GetDiaryById)
	diary.PATCH("/:id", r.UpdateDiary)
	diary.DELETE("/:id", r.DeleteDiary)
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Timeout())
	routerGroup := r.router.Group("/api/v1")
	MountDiary(routerGroup, r)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
