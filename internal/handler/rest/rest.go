package rest

import (
	"errors"
	"fmt"
	"memorabilia/internal/service"
	"memorabilia/pkg/bcrypt"
	"memorabilia/pkg/jwt"
	"memorabilia/pkg/middleware"
	"memorabilia/pkg/response"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
	jwt        jwt.Interface
	bcrypt     bcrypt.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface, jwt jwt.Interface, bcrypt bcrypt.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
		jwt:        jwt,
		bcrypt:     bcrypt,
	}
}

func MountDiary(routerGroup *gin.RouterGroup, r *Rest) {
	diary := routerGroup.Group("/diaries", r.middleware.AuthenticateUser)
	MountDiaryPicture(diary, r)
	diary.POST("", r.CreateDiary)
	diary.GET("/", r.GetDiary)
	diary.GET("/:diaryId", r.GetDiaryById)
	diary.PATCH("/:diaryId", r.UpdateDiary)
	diary.DELETE("/:diaryId", r.DeleteDiary)
}

func MountDiaryPicture(routerGroup *gin.RouterGroup, r *Rest) {
	diaryPicture := routerGroup.Group("/:diaryId/pictures")
	diaryPicture.POST("", r.AddDiaryPicture)
	diaryPicture.DELETE("/:pictureId", r.DeleteDiaryPicture)
}

func MountUser(routerGroup *gin.RouterGroup, r *Rest) {
	user := routerGroup.Group("/users")
	user.POST("/register", r.Register)
	user.POST("/login", r.Login)
	user.PUT("/profile-picture", r.middleware.AuthenticateUser, r.UploadProfilePicture)
	user.PATCH("", r.middleware.AuthenticateUser, r.UpdateProfile)
	user.GET("", r.middleware.AuthenticateUser, r.GetLoginUser)
}

func MountPoeple(routerGroup *gin.RouterGroup, r *Rest) {
	people := routerGroup.Group("/peoples")
	people.POST("", r.middleware.AuthenticateUser, r.CreatePeople)
	people.PATCH("/:peopleId", r.middleware.AuthenticateUser, r.UpdatePeople)
	people.DELETE("/:peopleId", r.middleware.AuthenticateUser, r.DeletePeople)
}

func (r *Rest) MountEndpoint() {
	r.router.Use(r.middleware.Timeout())

	r.router.NoRoute(func(ctx *gin.Context) {
		response.Error(ctx, http.StatusNotFound, "not found", errors.New("page not found"))
	})

	routerGroup := r.router.Group("/api/v1")
	MountDiary(routerGroup, r)
	MountUser(routerGroup, r)
	MountPoeple(routerGroup, r)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
