// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: state.sql

package database

import (
	"context"
)

const createState = `-- name: CreateState :one
INSERT INTO state(
    name
)VALUES(
    $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreateState(ctx context.Context, name string) (State, error) {
	row := q.db.QueryRowContext(ctx, createState, name)
	var i State
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteState = `-- name: DeleteState :exec
DELETE FROM state
WHERE id = $1
`

func (q *Queries) DeleteState(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteState, id)
	return err
}

const getState = `-- name: GetState :one
SELECT id, name, created_at FROM state
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetState(ctx context.Context, id int64) (State, error) {
	row := q.db.QueryRowContext(ctx, getState, id)
	var i State
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getStateIdByName = `-- name: GetStateIdByName :one
SELECT id FROM state
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetStateIdByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getStateIdByName, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const listStates = `-- name: ListStates :many
SELECT id, name, created_at FROM state
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListStatesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStates(ctx context.Context, arg ListStatesParams) ([]State, error) {
	rows, err := q.db.QueryContext(ctx, listStates, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []State{}
	for rows.Next() {
		var i State
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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

const updateState = `-- name: UpdateState :one
UPDATE state
set name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateStateParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateState(ctx context.Context, arg UpdateStateParams) (State, error) {
	row := q.db.QueryRowContext(ctx, updateState, arg.ID, arg.Name)
	var i State
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
