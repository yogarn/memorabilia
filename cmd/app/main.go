package main

import (
	"memorabilia/internal/handler/rest"
	"memorabilia/internal/repository"
	"memorabilia/internal/service"
	"memorabilia/pkg/config"
	"memorabilia/pkg/database/mysql"
)

func main() {
	config.LoadEnvironment()
	db := mysql.ConnectDatabase()
	defer db.Close()
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	rest := rest.NewRest(service)
	rest.MountEndpoint()
	rest.Run()
}
