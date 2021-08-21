package service

import (
	"context"
	"database/sql"
	"github.com/SemmiDev/todo-app/exception"
	"github.com/SemmiDev/todo-app/helper"
	"github.com/SemmiDev/todo-app/model/domain"
	"github.com/SemmiDev/todo-app/model/web"
	"github.com/SemmiDev/todo-app/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TodoService interface {
	Create(ctx context.Context, request web.CreateToDoRequest) *web.ToDoResponse
	Update(ctx context.Context, request web.UpdateToDoRequest) *web.ToDoResponse
	UpdateStatus(ctx context.Context, todoId string) *web.ToDoResponse
	Delete(ctx context.Context, todoId string)
	FindById(ctx context.Context, todoId string) *web.ToDoResponse
	FindAll(ctx context.Context) []*web.ToDoResponse
}

type TodoServiceImpl struct {
	TodoRepository repository.ToDoRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewTodoService(todoRepo repository.ToDoRepository, DB *sql.DB, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepo,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request web.CreateToDoRequest) *web.ToDoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	starting, ends, duration, timeErr := helper.ExtractTodoTime(request.StartingAt, request.EndsAt)
	if timeErr != nil {
		panic(exception.NewTimeNotValidError(timeErr.Error()))
	}

	todo := domain.ToDo{
		Id:         uuid.NewString(),
		Task:       request.Task,
		StartingAt: starting,
		EndsAt:     ends,
		Duration:   duration,
		IsExpired:  0,
		Done:       0,
	}

	todoResponse := service.TodoRepository.Save(ctx, tx, todo)

	return helper.ToTodoResponse(todoResponse)
}

func (service *TodoServiceImpl) Update(ctx context.Context, request web.UpdateToDoRequest) *web.ToDoResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	starting, ends, duration, timeErr := helper.ExtractTodoTime(request.StartingAt, request.EndsAt)
	if timeErr != nil {
		panic(exception.NewTimeNotValidError(timeErr.Error()))
	}

	todo.Task = request.Task
	todo.StartingAt = starting
	todo.EndsAt = ends
	todo.Duration = duration
	todo.IsExpired = 0
	todo.Done = 0

	todoUpdateResponse := service.TodoRepository.Update(ctx, tx, todo)

	return helper.ToTodoResponse(todoUpdateResponse)
}

func (service *TodoServiceImpl) UpdateStatus(ctx context.Context, todoId string) *web.ToDoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	todoUpdateStatusResponse := service.TodoRepository.UpdateStatus(ctx, tx, todo)

	return helper.ToTodoResponse(todoUpdateStatusResponse)
}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}

func (service *TodoServiceImpl) FindById(ctx context.Context, todoId string) *web.ToDoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTodoResponse(todo)
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []*web.ToDoResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	todos := service.TodoRepository.FindAll(ctx, tx)

	return helper.ToTodoResponses(todos)
}
