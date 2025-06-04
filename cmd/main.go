package main

import (
	"log"

	todo "github.com/zenmaster911/shelfAPI"
	"github.com/zenmaster911/shelfAPI/pkg/handler"
	"github.com/zenmaster911/shelfAPI/pkg/repository"
	"github.com/zenmaster911/shelfAPI/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}

}
