-- name: CreateTeam :one
INSERT INTO team(
    name
)VALUES(
    $1
)
RETURNING *;