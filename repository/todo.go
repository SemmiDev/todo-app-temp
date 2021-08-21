package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/SemmiDev/todo-app/helper"
	"github.com/SemmiDev/todo-app/model/domain"
)

type ToDoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo
	Update(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo
	UpdateStatus(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo
	Delete(ctx context.Context, tx *sql.Tx, todo domain.ToDo)
	FindById(ctx context.Context, tx *sql.Tx, todoId string) (domain.ToDo, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ToDo
}

type ToDoRepositoryImpl struct{}

func (c *ToDoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo {
	_, err := tx.ExecContext(ctx, SaveTodoCmd,
		todo.Id,
		todo.Task,
		todo.StartingAt,
		todo.EndsAt,
		todo.Duration,
		todo.IsExpired,
		todo.Done)

	helper.PanicIfError(err)
	return todo
}

func (c *ToDoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo {
	_, err := tx.ExecContext(ctx, UpdateTodoCmd,
		todo.Task,
		todo.StartingAt,
		todo.EndsAt,
		todo.Duration,
		todo.IsExpired,
		todo.Done,
		todo.Id)

	helper.PanicIfError(err)
	return todo
}

func (c *ToDoRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo {
	_, err := tx.ExecContext(ctx, UpdateStatusTodoCmd, 1, todo.Id)
	helper.PanicIfError(err)

	todo.Done = 1
	return todo
}

func (c *ToDoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.ToDo) {
	_, err := tx.ExecContext(ctx, DeleteTodoCmd, todo.Id)
	helper.PanicIfError(err)
}

func (c *ToDoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId string) (domain.ToDo, error) {
	rows, err := tx.QueryContext(ctx, FindByIdQuery, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.ToDo{}
	if rows.Next() {
		err := rows.Scan(
			&todo.Id,
			&todo.Task,
			&todo.StartingAt,
			&todo.EndsAt,
			&todo.Duration,
			&todo.IsExpired,
			&todo.Done)

		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (c *ToDoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ToDo {
	rows, err := tx.QueryContext(ctx, FindAllQuery)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.ToDo
	for rows.Next() {
		todo := domain.ToDo{}
		err := rows.Scan(
			&todo.Id,
			&todo.Task,
			&todo.StartingAt,
			&todo.EndsAt,
			&todo.Duration,
			&todo.IsExpired,
			&todo.Done)

		helper.PanicIfError(err)
		todos = append(todos, todo)
	}

	return todos
}

func NewToDoRepository() ToDoRepository {
	return &ToDoRepositoryImpl{}
}
