-- name: CreateProject :one
INSERT INTO project(
    name
)VALUES(
    $1
)
RETURNING *;