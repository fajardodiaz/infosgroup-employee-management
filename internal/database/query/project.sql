-- name: CreateProject :one
INSERT INTO project(
    name
)VALUES(
    $1
)
RETURNING *;

-- name: GetProject :one
SELECT * FROM project
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProject :one
UPDATE project
set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1;