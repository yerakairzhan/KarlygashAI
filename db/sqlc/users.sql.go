// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (userid, username)
VALUES ($1, $2)
    RETURNING userid, username, created_at
`

type CreateUserParams struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Userid, arg.Username)
	var i User
	err := row.Scan(&i.Userid, &i.Username, &i.CreatedAt)
	return i, err
}
