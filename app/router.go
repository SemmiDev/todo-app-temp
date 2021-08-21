package app

import (
	"github.com/SemmiDev/todo-app/controller"
	"github.com/SemmiDev/todo-app/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoController controller.ToDoController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/todos", todoController.Create)
	router.GET("/api/todos", todoController.FindAll)
	router.GET("/api/todos/:todoId", todoController.FindById)
	router.PUT("/api/todos/:todoId", todoController.Update)
	router.PATCH("/api/todos/:todoId", todoController.UpdateStatus)
	router.DELETE("/api/todos/:todoId", todoController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
