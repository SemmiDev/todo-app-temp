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
	SQL := `insert into todos(id, task, starting_at, ends_at, duration, is_expired, done) values (?,?,?,?,?,?,?)`
	_, err := tx.ExecContext(ctx, SQL, todo.Id, todo.Task, todo.StartingAt, todo.EndsAt, todo.Duration, todo.IsExpired, todo.Done)
	helper.PanicIfError(err)
	return todo
}

func (c *ToDoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo {
	SQL := `update todos set task = ?, starting_at = ?, ends_at = ?, duration = ?, is_expired = ?, done = ? where id = ?`
	_, err := tx.ExecContext(ctx, SQL, todo.Task, todo.StartingAt, todo.EndsAt, todo.Duration, todo.IsExpired, todo.Done, todo.Id)
	helper.PanicIfError(err)
	return todo
}

func (c *ToDoRepositoryImpl) UpdateStatus(ctx context.Context, tx *sql.Tx, todo domain.ToDo) domain.ToDo {
	SQL := `update todos set done = ? where id = ?`
	_, err := tx.ExecContext(ctx, SQL, 1, todo.Id)
	helper.PanicIfError(err)

	todo.Done = 1
	return todo
}

func (c *ToDoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.ToDo) {
	SQL := "delete from todos where id = ?"
	_, err := tx.ExecContext(ctx, SQL, todo.Id)
	helper.PanicIfError(err)
}

func (c *ToDoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId string) (domain.ToDo, error) {
	SQL := `select id, task, starting_at, ends_at, duration, is_expired, done from todos where id = ?`
	rows, err := tx.QueryContext(ctx, SQL, todoId)
	helper.PanicIfError(err)
	defer rows.Close()

	todo := domain.ToDo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Task, &todo.StartingAt, &todo.EndsAt, &todo.Duration, &todo.IsExpired, &todo.Done)
		helper.PanicIfError(err)
		return todo, nil
	} else {
		return todo, errors.New("todo is not found")
	}
}

func (c *ToDoRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ToDo {
	SQL := `select id, task, starting_at, ends_at, duration, is_expired, done from todos`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var todos []domain.ToDo
	for rows.Next() {
		todo := domain.ToDo{}
		err := rows.Scan(&todo.Id, &todo.Task, &todo.StartingAt, &todo.EndsAt, &todo.Duration, &todo.IsExpired, &todo.Done)
		helper.PanicIfError(err)
		todos = append(todos, todo)
	}
	return todos
}

func NewToDoRepository() ToDoRepository {
	return &ToDoRepositoryImpl{}
}
