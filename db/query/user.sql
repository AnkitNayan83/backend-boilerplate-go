-- name: CreateUser :one
INSERT INTO users (name, email, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, email, created_at, updated_at;


-- name: GetUserByID :one
SELECT id, name, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: UpdateUserByID :one
UPDATE users
SET
    name = COALESCE(sqlc.narg(name), name),
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password)
WHERE id = sqlc.narg(id)
RETURNING *;


-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1;