-- name: GetPostByID :one
SELECT id, title, content, author_id, created_at, updated_at
FROM post
WHERE id = $1;

-- name: ListPosts :many
SELECT id, title, content, author_id, created_at, updated_at
FROM post
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreatePost :one
INSERT INTO post (title, content, author_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdatePostByID :one
UPDATE post
SET
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    author_id = COALESCE(sqlc.narg(author_id), author_id),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.narg(id)
RETURNING *;

-- name: DeletePostByID :exec
DELETE FROM post
WHERE id = $1;