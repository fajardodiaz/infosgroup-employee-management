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

func CreateRandomPosition(t *testing.T) database.Position {
	randomName := utils.RandomString(6)
	position1, err := testQueries.CreatePosition(context.Background(), randomName)
	require.NoError(t, err)
	require.NotEmpty(t, position1)
	require.Equal(t, position1.Name, randomName)
	require.NotZero(t, position1.ID)
	require.NotZero(t, position1.Name)

	return position1
}

func TestCreatePosition(t *testing.T) {
	CreateRandomPosition(t)
}

func TestGetPosition(t *testing.T) {
	position1 := CreateRandomPosition(t)

	position2, err := testQueries.GetPosition(context.Background(), position1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, position2)
	require.Equal(t, position1, position2)
	require.Equal(t, position1.ID, position2.ID)
	require.Equal(t, position1.Name, position2.Name)
	require.WithinDuration(t, position1.CreatedAt, position2.CreatedAt, time.Second)
}

func TestGetPositonIdByName(t *testing.T) {
	position1 := CreateRandomPosition(t)

	position2, err := testQueries.GetPositionIdByName(context.Background(), position1.Name)

	require.NoError(t, err)
	require.NotEmpty(t, position2)
	require.Equal(t, position1.ID, position2)
}

func TestListPositions(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomPosition(t)
	}

	args := database.ListPositionsParams{
		Limit:  5,
		Offset: 5,
	}

	positions, err := testQueries.ListPositions(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, positions, 5)

	for _, position := range positions {
		require.NotEmpty(t, position)
	}

}

func TestUpdatePosition(t *testing.T) {
	position1 := CreateRandomPosition(t)

	args := database.UpdatePositionParams{
		ID:   position1.ID,
		Name: utils.RandomString(8),
	}

	position2, err := testQueries.UpdatePosition(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, position1.ID, args.ID)
	require.Equal(t, position2.Name, args.Name)

	require.NotEqual(t, position1.Name, position2.Name)

}

func TestDeletePosition(t *testing.T) {
	position1 := CreateRandomPosition(t)

	testQueries.DeletePosition(context.Background(), position1.ID)

	position2, err := testQueries.GetPosition(context.Background(), position1.ID)

	require.Error(t, err)
	require.Empty(t, position2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
