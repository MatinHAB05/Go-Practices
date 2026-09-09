-- name: CreateUser :execresult
INSERT INTO users (name, email)
VALUES (?, ?);

-- name: GetUser :one
SELECT id, name
FROM users
WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, email
FROM users
ORDER BY id;
