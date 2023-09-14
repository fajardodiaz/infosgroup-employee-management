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

func CreateRandomTeam(t *testing.T) database.Team {
	randomName := utils.RandomString(6)
	team1, err := testQueries.CreateTeam(context.Background(), randomName)
	require.NoError(t, err)
	require.NotEmpty(t, team1)
	require.Equal(t, team1.Name, randomName)
	require.NotZero(t, team1.ID)
	require.NotZero(t, team1.Name)

	return team1
}

func TestCreateTeam(t *testing.T) {
	CreateRandomTeam(t)
}

func TestGetTeam(t *testing.T) {
	team1 := CreateRandomTeam(t)

	team2, err := testQueries.GetTeam(context.Background(), team1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, team2)
	require.Equal(t, team1, team2)
	require.Equal(t, team1.ID, team2.ID)
	require.Equal(t, team1.Name, team2.Name)
	require.WithinDuration(t, team1.CreatedAt, team2.CreatedAt, time.Second)
}

func TestGetTeamIdByName(t *testing.T) {
	team1 := CreateRandomTeam(t)

	team2, err := testQueries.GetTeamIdByName(context.Background(), team1.Name)

	require.NoError(t, err)
	require.NotEmpty(t, team2)
	require.Equal(t, team1.ID, team2)
}

func TestListTeams(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTeam(t)
	}

	args := database.ListTeamsParams{
		Limit:  5,
		Offset: 5,
	}

	teams, err := testQueries.ListTeams(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, teams, 5)

	for _, team := range teams {
		require.NotEmpty(t, team)
	}

}

func TestUpdateTeam(t *testing.T) {
	team1 := CreateRandomTeam(t)

	args := database.UpdateTeamParams{
		ID:   team1.ID,
		Name: utils.RandomString(8),
	}

	team2, err := testQueries.UpdateTeam(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, team1.ID, args.ID)
	require.Equal(t, team2.Name, args.Name)

	require.NotEqual(t, team1.Name, team2.Name)

}

func TestDeleteTeam(t *testing.T) {
	team1 := CreateRandomTeam(t)

	testQueries.DeleteTeam(context.Background(), team1.ID)

	team2, err := testQueries.GetTeam(context.Background(), team1.ID)

	require.Error(t, err)
	require.Empty(t, team2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
