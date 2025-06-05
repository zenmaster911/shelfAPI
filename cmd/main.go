package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/zenmaster911/shelfAPI"
	"github.com/zenmaster911/shelfAPI/pkg/handler"
	"github.com/zenmaster911/shelfAPI/pkg/repository"
	"github.com/zenmaster911/shelfAPI/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error in initialization configs: %v", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
