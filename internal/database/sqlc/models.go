// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"
)

type Employee struct {
	ID                int64          `json:"id"`
	EmployeeCod       string         `json:"employee_cod"`
	FullName          string         `json:"full_name"`
	Birth             time.Time      `json:"birth"`
	IngressDate       time.Time      `json:"ingress_date"`
	EndEvaluationDate time.Time      `json:"end_evaluation_date"`
	Phone             sql.NullString `json:"phone"`
	Gender            sql.NullString `json:"gender"`
	CreatedAt         sql.NullTime   `json:"created_at"`
	StateID           sql.NullInt32  `json:"state_id"`
	PositionID        sql.NullInt32  `json:"position_id"`
	TeamID            sql.NullInt32  `json:"team_id"`
}

type EmployeeProject struct {
	EmployeeID int64 `json:"employee_id"`
	ProjectID  int64 `json:"project_id"`
}

type Position struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type State struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Team struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}