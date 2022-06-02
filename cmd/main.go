package main

import (
	"log"

	"github.com/emPeeGee/todo-go"
	"github.com/emPeeGee/todo-go/pkg/handler"
	"github.com/emPeeGee/todo-go/pkg/repository"
	"github.com/emPeeGee/todo-go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running server %s", err.Error())
	}
}
