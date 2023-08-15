package tests

import (
	"context"
	"testing"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/stretchr/testify/require"
)

func TestAssignEmployeeToProject(t *testing.T) {
	project1 := CreateRandomProject(t)
	employee1 := CreateRandomEmployee(t)

	args := database.AssignEmployeeToProjectParams{
		EmployeeID: employee1.ID,
		ProjectID:  project1.ID,
	}

	assignment, err := testQueries.AssignEmployeeToProject(context.Background(), args)

	require.NoError(t, err)

	require.NotEmpty(t, assignment)
	require.Equal(t, assignment.EmployeeID, employee1.ID)
	require.Equal(t, assignment.ProjectID, project1.ID)
}
