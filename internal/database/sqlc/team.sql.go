// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: team.sql

package database

import (
	"context"
)

const createTeam = `-- name: CreateTeam :one
INSERT INTO team(
    name
)VALUES(
    $1
)
RETURNING id, name
`

func (q *Queries) CreateTeam(ctx context.Context, name string) (Team, error) {
	row := q.db.QueryRowContext(ctx, createTeam, name)
	var i Team
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
