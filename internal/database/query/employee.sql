-- name: CreateEmployee :one
INSERT INTO employee (
    employee_cod,
    full_name,
    birth,
    ingress_date,
    end_evaluation_date,
    phone,
    gender,
    state_id,
    position_id,
    team_id
)
VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetEmployee :one
SELECT * FROM employee
WHERE id = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employee
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEmployee :one
UPDATE employee set
employee_cod = $2,
full_name = $3,
birth = $4,
ingress_date = $5,
end_evaluation_date = $6,
phone = $7,
gender = $8,
state_id = $9,
position_id = $10,
team_id = $11
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE id = $1;