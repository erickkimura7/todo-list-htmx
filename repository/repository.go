package repository

import (
	"context"
	"database/sql"
	"errors"
	"todo-list/database"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type TodosDatabase struct {
	query *database.Queries
}

type DatabaseImpl interface {
	CreateTodo(title string) error
	GetAllTodos() ([]database.Todo, error)
	DeleteById(id string) error
	UpdateStatus(id string) (database.Todo, error)
}

func New(db *sql.DB) *TodosDatabase {
	return &TodosDatabase{
		query: database.New(db),
	}
}

func (d *TodosDatabase) CreateTodo(title string) (database.Todo, error) {
	todo, err := d.query.CreateTodos(context.Background(), database.CreateTodosParams{
		ID:     uuid.NewString(),
		Title:  title,
		IsDone: false,
	})

	return todo, err
}

func (d *TodosDatabase) GetAllTodos() ([]database.Todo, error) {
	todos, err := d.query.GetAllTodos(context.Background())
	if err != nil {
		return nil, errors.Join(errors.New("get all todos"), err)
	}

	return todos, nil
}

func (d *TodosDatabase) DeleteById(id string) error {
	return d.query.DeleteTodo(context.Background(), id)
}

func (d *TodosDatabase) UpdateStatus(id string) (database.Todo, error) {
	todo, err := d.query.GetTodos(context.Background(), id)
	if err != nil {
		return database.Todo{}, err
	}

	newTodo, err := d.query.UpdateTodo(context.Background(), database.UpdateTodoParams{
		IsDone: !todo.IsDone,
		ID:     id,
	})

	return newTodo, err
}
