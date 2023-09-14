// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: position.sql

package database

import (
	"context"
)

const createPosition = `-- name: CreatePosition :one
INSERT INTO position(
    name
)VALUES(
    $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreatePosition(ctx context.Context, name string) (Position, error) {
	row := q.db.QueryRowContext(ctx, createPosition, name)
	var i Position
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deletePosition = `-- name: DeletePosition :exec
DELETE FROM position
WHERE id = $1
`

func (q *Queries) DeletePosition(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePosition, id)
	return err
}

const getPosition = `-- name: GetPosition :one
SELECT id, name, created_at FROM position
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPosition(ctx context.Context, id int64) (Position, error) {
	row := q.db.QueryRowContext(ctx, getPosition, id)
	var i Position
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const getPositionIdByName = `-- name: GetPositionIdByName :one
SELECT id FROM position
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPositionIdByName(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPositionIdByName, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const listPositions = `-- name: ListPositions :many
SELECT id, name, created_at FROM position
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPositionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPositions(ctx context.Context, arg ListPositionsParams) ([]Position, error) {
	rows, err := q.db.QueryContext(ctx, listPositions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Position{}
	for rows.Next() {
		var i Position
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

const updatePosition = `-- name: UpdatePosition :one
UPDATE position
set name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdatePositionParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePosition(ctx context.Context, arg UpdatePositionParams) (Position, error) {
	row := q.db.QueryRowContext(ctx, updatePosition, arg.ID, arg.Name)
	var i Position
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
