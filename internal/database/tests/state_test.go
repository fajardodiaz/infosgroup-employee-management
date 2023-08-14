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

func CreateRandomState(t *testing.T) database.State {
	randomName := utils.RandomString(6)
	state1, err := testQueries.CreateState(context.Background(), randomName)
	require.NoError(t, err)
	require.NotEmpty(t, state1)
	require.Equal(t, state1.Name, randomName)
	require.NotZero(t, state1.ID)
	require.NotZero(t, state1.Name)

	return state1
}

func TestCreateState(t *testing.T) {
	CreateRandomState(t)
}

func TestGetState(t *testing.T) {
	state1 := CreateRandomState(t)

	state2, err := testQueries.GetState(context.Background(), state1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, state2)
	require.Equal(t, state1, state2)
	require.Equal(t, state1.ID, state2.ID)
	require.Equal(t, state1.Name, state2.Name)
	require.WithinDuration(t, state1.CreatedAt, state2.CreatedAt, time.Second)

}

func TestListStates(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomState(t)
	}

	args := database.ListStatesParams{
		Limit:  5,
		Offset: 5,
	}

	states, err := testQueries.ListStates(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, states, 5)

	for _, state := range states {
		require.NotEmpty(t, state)
	}

}

func TestUpdateState(t *testing.T) {
	state1 := CreateRandomState(t)

	args := database.UpdateStateParams{
		ID:   state1.ID,
		Name: utils.RandomString(8),
	}

	state2, err := testQueries.UpdateState(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, state1.ID, args.ID)
	require.Equal(t, state2.Name, args.Name)

	require.NotEqual(t, state1.Name, state2.Name)

}

func TestDeleteState(t *testing.T) {
	state1 := CreateRandomState(t)

	testQueries.DeleteState(context.Background(), state1.ID)

	state2, err := testQueries.GetState(context.Background(), state1.ID)

	require.Error(t, err)
	require.Empty(t, state2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
