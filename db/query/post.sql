-- name: GetPostByID :one
SELECT id, title, content, author_id, created_at, updated_at
FROM post
WHERE id = $1;