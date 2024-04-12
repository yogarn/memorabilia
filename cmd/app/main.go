package main

import (
	"memorabilia/internal/handler/rest"
	"memorabilia/internal/repository"
	"memorabilia/internal/service"
	"memorabilia/pkg/config"
	"memorabilia/pkg/database/mysql"
	"memorabilia/pkg/middleware"
)

func main() {
	config.LoadEnvironment()
	db := mysql.ConnectDatabase()
	defer db.Close()
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	middleware := middleware.Init()
	rest := rest.NewRest(service, middleware)
	rest.MountEndpoint()
	rest.Run()
}
