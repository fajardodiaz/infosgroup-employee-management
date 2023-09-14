// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: project.sql

package database

import (
	"context"
)

const createProject = `-- name: CreateProject :one
INSERT INTO project(
    name
)VALUES(
    $1
)
RETURNING id, name, created_at
`

func (q *Queries) CreateProject(ctx context.Context, name string) (Project, error) {
	row := q.db.QueryRowContext(ctx, createProject, name)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const getProject = `-- name: GetProject :one
SELECT id, name, created_at FROM project
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProject(ctx context.Context, id int64) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listProjects = `-- name: ListProjects :many
SELECT id, name, created_at FROM project
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProjectsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProjects(ctx context.Context, arg ListProjectsParams) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, listProjects, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Project{}
	for rows.Next() {
		var i Project
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

const updateProject = `-- name: UpdateProject :one
UPDATE project
set name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateProjectParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, updateProject, arg.ID, arg.Name)
	var i Project
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
