package tests

import (
	"context"
	"database/sql"
	"testing"
	"time"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/fajardodiaz/infosgroup-employee-management/internal/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomProject(t *testing.T) database.Project {
	randomName := utils.RandomString(6)
	project1, err := testQueries.CreateProject(context.Background(), randomName)
	require.NoError(t, err)
	require.NotEmpty(t, project1)
	require.Equal(t, project1.Name, randomName)
	require.NotZero(t, project1.ID)
	require.NotZero(t, project1.Name)

	return project1
}

func TestCreateProject(t *testing.T) {
	CreateRandomProject(t)
}

func TestGetProject(t *testing.T) {
	project1 := CreateRandomProject(t)

	project2, err := testQueries.GetProject(context.Background(), project1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, project2)
	require.Equal(t, project1, project2)
	require.Equal(t, project1.ID, project2.ID)
	require.Equal(t, project1.Name, project2.Name)
	require.WithinDuration(t, project1.CreatedAt, project2.CreatedAt, time.Second)

}

func TestListProjects(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomProject(t)
	}

	args := database.ListProjectsParams{
		Limit:  5,
		Offset: 5,
	}

	projects, err := testQueries.ListProjects(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, projects, 5)

	for _, project := range projects {
		require.NotEmpty(t, project)
	}

}

func TestUpdateProject(t *testing.T) {
	project1 := CreateRandomProject(t)

	args := database.UpdateProjectParams{
		ID:   project1.ID,
		Name: utils.RandomString(8),
	}

	project2, err := testQueries.UpdateProject(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, project1.ID, args.ID)
	require.Equal(t, project2.Name, args.Name)

	require.NotEqual(t, project1.Name, project2.Name)

}

func TestDeleteProject(t *testing.T) {
	project1 := CreateRandomProject(t)

	testQueries.DeleteProject(context.Background(), project1.ID)

	project2, err := testQueries.GetProject(context.Background(), project1.ID)

	require.Error(t, err)
	require.Empty(t, project2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
