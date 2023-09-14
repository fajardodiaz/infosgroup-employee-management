-- name: CreateTeam :one
INSERT INTO team(
    name
)VALUES(
    $1
)
RETURNING *;

-- name: GetTeam :one
SELECT * FROM team
WHERE id = $1 LIMIT 1;

-- name: GetTeamIdByName :one
SELECT id FROM team
WHERE name = $1 LIMIT 1;

-- name: ListTeams :many
SELECT * FROM team
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTeam :one
UPDATE team
set name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTeam :exec
DELETE FROM team
WHERE id = $1;