-- name: CreatePosition :one
INSERT INTO position(
    name
)VALUES(
    $1
)
RETURNING *;

-- name: GetPosition :one
SELECT * FROM position
WHERE id = $1 LIMIT 1;

-- name: GetPositionIdByName :one
SELECT id FROM position
WHERE name = $1 LIMIT 1;

-- name: ListPositions :many
SELECT * FROM position
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePosition :one
UPDATE position
set name = $2
WHERE id = $1
RETURNING *;

-- name: DeletePosition :exec
DELETE FROM position
WHERE id = $1;