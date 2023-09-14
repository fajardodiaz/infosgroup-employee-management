-- name: CreateState :one
INSERT INTO state(
    name
)VALUES(
    $1
)
RETURNING *;

-- name: GetState :one
SELECT * FROM state
WHERE id = $1 LIMIT 1;

-- name: GetStateIdByName :one
SELECT id FROM state
WHERE name = $1 LIMIT 1;

-- name: ListStates :many
SELECT * FROM state
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateState :one
UPDATE state
set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteState :exec
DELETE FROM state
WHERE id = $1;