-- name: AssignEmployeeToProject :one
INSERT INTO employee_project(
    employee_id,
    project_id
) VALUES(
    $1, $2
)
RETURNING *;
