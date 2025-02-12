-- name: CreateUser :one
INSERT INTO users (userid, username)
VALUES ($1, $2)
    RETURNING userid, username, created_at;