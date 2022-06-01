package main

import (
	"log"

	"github.com/emPeeGee/todo-go"
	"github.com/emPeeGee/todo-go/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("Error occurred while running server %s", err.Error())
	}
}
