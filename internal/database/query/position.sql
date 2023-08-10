-- name: CreatePosition :one
INSERT INTO position(
    name
)VALUES(
    $1
)
RETURNING *;