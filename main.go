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
	"net/http"
)

func main() {
	db := app.NewDB()
	validation := validator.New()
	todoRepository := repository.NewToDoRepository()
	todoService := service.NewTodoService(todoRepository, db, validation)
	todoController := controller.NewToDoController(todoService)
	router := app.NewRouter(todoController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
