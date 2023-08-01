// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package database

import (
	"context"
)

const createTodos = `-- name: CreateTodos :one
INSERT INTO todos (
  id, title, is_done
) VALUES (
 ?, ?, ? 
)
RETURNING id, title, is_done
`

type CreateTodosParams struct {
	ID     string
	Title  string
	IsDone bool
}

func (q *Queries) CreateTodos(ctx context.Context, arg CreateTodosParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodos, arg.ID, arg.Title, arg.IsDone)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.IsDone)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getAllTodos = `-- name: GetAllTodos :many
SELECT id, title, is_done FROM todos
`

func (q *Queries) GetAllTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Title, &i.IsDone); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodos = `-- name: GetTodos :one
SELECT id, title, is_done FROM todos
WHERE id = ?
`

func (q *Queries) GetTodos(ctx context.Context, id string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodos, id)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.IsDone)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET is_done = ?
WHERE id = ?
RETURNING id, title, is_done
`

type UpdateTodoParams struct {
	IsDone bool
	ID     string
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.IsDone, arg.ID)
	var i Todo
	err := row.Scan(&i.ID, &i.Title, &i.IsDone)
	return i, err
}
