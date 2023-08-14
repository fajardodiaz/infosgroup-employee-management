package tests

import (
	"context"
	"fmt"
	"testing"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/fajardodiaz/infosgroup-employee-management/internal/utils"

	"github.com/stretchr/testify/require"
)

func CreateRandomTeam(t *testing.T) database.Team {
	randomName := utils.RandomString(6)
	team1, err := testQueries.CreateTeam(context.Background(), randomName)
	require.NoError(t, err)
	require.NotEmpty(t, team1)
	fmt.Println(err)

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

}

// func TestListTeams(t *testing.T)  {}
// func TestUpdateTeam(t *testing.T) {}
// func TestDeleteTeam(t *testing.T) {}
