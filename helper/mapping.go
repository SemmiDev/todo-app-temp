package helper

import (
	"github.com/SemmiDev/todo-app/model/domain"
	"github.com/SemmiDev/todo-app/model/web"
)

func ToTodoResponse(todo domain.ToDo) *web.ToDoResponse {
	isExpired := IsTaskExpired(todo.EndsAt)
	expired := 0
	if isExpired {
		expired = 1
	}

	return &web.ToDoResponse{
		Id:         todo.Id,
		Task:       todo.Task,
		StartingAt: todo.StartingAt,
		EndsAt:     todo.EndsAt,
		Duration:   todo.Duration,
		Expired:    expired,
		Done:       todo.Done,
	}
}

func ToTodoResponses(todos []domain.ToDo) []*web.ToDoResponse {
	var todosResponse []*web.ToDoResponse
	for _, todo := range todos {
		todosResponse = append(todosResponse, ToTodoResponse(todo))
	}
	return todosResponse
}
