package main

import (
	"memorabilia/internal/handler/rest"
	"memorabilia/internal/repository"
	"memorabilia/internal/service"
	"memorabilia/pkg/bcrypt"
	"memorabilia/pkg/config"
	"memorabilia/pkg/database/mysql"
	"memorabilia/pkg/jwt"
	"memorabilia/pkg/middleware"
	"memorabilia/pkg/supabase"
)

func main() {
	config.LoadEnvironment()
	db := mysql.ConnectDatabase()
	defer db.Close()

	jwt := jwt.Init()
	bcrypt := bcrypt.Init()
	supabase := supabase.Init()

	repository := repository.NewRepository(db)
	service := service.NewService(repository, bcrypt, jwt, supabase)
	middleware := middleware.Init(jwt, service)
	rest := rest.NewRest(service, middleware, jwt, bcrypt)
	rest.MountEndpoint()
	rest.Run()
}
