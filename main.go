package main

import (
	"github.com/SemmiDev/todo-app/app"
	"github.com/SemmiDev/todo-app/controller"
	"github.com/SemmiDev/todo-app/helper"
	"github.com/SemmiDev/todo-app/middleware"
	"github.com/SemmiDev/todo-app/repository"
	"github.com/SemmiDev/todo-app/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	db := app.NewDB()
	validation := validator.New()
	repository := repository.NewToDoRepository()
	service := service.NewTodoService(repository, db, validation)
	controller := controller.NewToDoController(service)
	router := app.NewRouter(controller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("running :: http://localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
