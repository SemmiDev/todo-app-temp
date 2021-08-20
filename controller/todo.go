package controller

import (
	"github.com/SemmiDev/todo-app/helper"
	"github.com/SemmiDev/todo-app/model/web"
	"github.com/SemmiDev/todo-app/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type ToDoController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ToDoControllerImpl struct {
	ToDoService service.TodoService
}

func NewToDoController(todoService service.TodoService) ToDoController {
	return &ToDoControllerImpl{
		ToDoService: todoService,
	}
}

func (controller *ToDoControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CreateToDoRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	todoResponse := controller.ToDoService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoUpdateReq := web.UpdateToDoRequest{}
	helper.ReadFromRequestBody(request, &todoUpdateReq)

	todoId := params.ByName("todoId")
	todoUpdateReq.Id = todoId
	log.Println(todoUpdateReq)

	todoResponse := controller.ToDoService.Update(request.Context(), todoUpdateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) UpdateStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")

	todoResponse := controller.ToDoService.UpdateStatus(request.Context(), todoId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")

	controller.ToDoService.Delete(request.Context(), todoId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")

	todoResponse := controller.ToDoService.FindById(request.Context(), todoId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ToDoControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponses := controller.ToDoService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
