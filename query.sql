-- name: GetAllTodos :many
SELECT * FROM todos;

-- name: GetTodos :one
SELECT * FROM todos
WHERE id = ?;

-- name: CreateTodos :one
INSERT INTO todos (
  id, title, is_done
) VALUES (
 ?, ?, ? 
)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos SET is_done = ?
WHERE id = ?
RETURNING *; 

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
