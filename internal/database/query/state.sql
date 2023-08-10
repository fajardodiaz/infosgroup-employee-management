-- name: CreateState :one
INSERT INTO state(
    name
)VALUES(
    $1
)
RETURNING *;