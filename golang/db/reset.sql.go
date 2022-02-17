// Code generated by sqlc. DO NOT EDIT.
// source: reset.sql

package db

import (
	"context"
)

const createReset = `-- name: CreateReset :one
INSERT INTO resets (user_id, code) VALUES ($1, $2) RETURNING user_id, code, created_at
`

type CreateResetParams struct {
	UserID int64  `json:"user_id"`
	Code   string `json:"code"`
}

func (q *Queries) CreateReset(ctx context.Context, arg CreateResetParams) (Reset, error) {
	row := q.db.QueryRowContext(ctx, createReset, arg.UserID, arg.Code)
	var i Reset
	err := row.Scan(&i.UserID, &i.Code, &i.CreatedAt)
	return i, err
}

const deleteResetsForUser = `-- name: DeleteResetsForUser :exec
DELETE FROM resets WHERE user_id = $1
`

func (q *Queries) DeleteResetsForUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteResetsForUser, userID)
	return err
}

const findResetByCode = `-- name: FindResetByCode :one
SELECT user_id, code, created_at FROM resets WHERE code = $1 LIMIT 1
`

func (q *Queries) FindResetByCode(ctx context.Context, code string) (Reset, error) {
	row := q.db.QueryRowContext(ctx, findResetByCode, code)
	var i Reset
	err := row.Scan(&i.UserID, &i.Code, &i.CreatedAt)
	return i, err
}
