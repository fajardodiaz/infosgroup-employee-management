-- name: AssignEmployeeToProject :one
INSERT INTO employee_project(
    employee_id,
    project_id
) VALUES(
    $1, $2
)
RETURNING *;

-- name: RemoveEmployeeProject :exec
DELETE FROM employee_project 
WHERE employee_id = $1 AND project_id = $2;