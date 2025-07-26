-- name: CreateUser :one
INSERT INTO users (user_id, username, name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1;

-- name: VerifyUser :one
SELECT * FROM users WHERE email = $1 AND password = $2;