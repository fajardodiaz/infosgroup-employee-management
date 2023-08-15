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

func CreateRandomEmployee(t *testing.T) database.Employee {
	randomCod := utils.RandomString(5)
	randomName := utils.RandomString(20)
	randomBirth := utils.RandomBirthDate()
	randomIngressDate := utils.RandomIngressDate()
	randomEndEvaluationDate := utils.Add3MonthsToDate(randomIngressDate)
	randomPhone := utils.RandonPhone(60000000, 69999999)
	randomGender := utils.RandomGender()
	randomState := CreateRandomState(t)
	randomPosition := CreateRandomPosition(t)
	randomTeam := CreateRandomTeam(t)

	args := database.CreateEmployeeParams{
		EmployeeCod:       randomCod,
		FullName:          randomName,
		Birth:             randomBirth,
		IngressDate:       randomIngressDate,
		EndEvaluationDate: randomEndEvaluationDate,
		Phone:             sql.NullString{String: randomPhone, Valid: true},
		Gender:            sql.NullString{String: randomGender, Valid: true},
		StateID:           sql.NullInt32{Int32: int32(randomState.ID), Valid: true},
		PositionID:        sql.NullInt32{Int32: int32(randomPosition.ID), Valid: true},
		TeamID:            sql.NullInt32{Int32: int32(randomTeam.ID), Valid: true},
	}

	employee1, err := testQueries.CreateEmployee(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, employee1)
	require.Equal(t, employee1.FullName, randomName)
	require.NotZero(t, employee1.ID)
	require.NotZero(t, employee1.FullName)

	return employee1
}

func TestCreateEmployee(t *testing.T) {
	CreateRandomEmployee(t)
}

func TestGetEmployee(t *testing.T) {
	employee1 := CreateRandomEmployee(t)

	employee2, err := testQueries.GetEmployee(context.Background(), employee1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, employee2)
	require.Equal(t, employee1, employee2)
	require.Equal(t, employee1.ID, employee2.ID)
	require.Equal(t, employee1.FullName, employee2.FullName)
	require.Equal(t, employee1.Gender, employee2.Gender)

	require.WithinDuration(t, employee1.CreatedAt, employee2.CreatedAt, time.Second)
}

func TestListEmployees(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEmployee(t)
	}

	args := database.ListEmployeesParams{
		Limit:  5,
		Offset: 5,
	}

	employees, err := testQueries.ListEmployees(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, employees, 5)

	for _, position := range employees {
		require.NotEmpty(t, position)
	}

}

func TestUpdateEmployee(t *testing.T) {
	employee1 := CreateRandomEmployee(t)

	ingress := utils.RandomIngressDate()
	args := database.UpdateEmployeeParams{
		ID:                employee1.ID,
		EmployeeCod:       utils.RandomString(6),
		FullName:          utils.RandomString(30),
		Birth:             utils.RandomBirthDate(),
		IngressDate:       ingress,
		EndEvaluationDate: utils.Add3MonthsToDate(ingress),
		Phone:             sql.NullString{String: utils.RandonPhone(60000000, 69999999), Valid: true},
		Gender:            sql.NullString{String: utils.RandomGender(), Valid: true},
		StateID:           sql.NullInt32{Int32: int32(CreateRandomState(t).ID), Valid: true},
		PositionID:        sql.NullInt32{Int32: int32(CreateRandomPosition(t).ID), Valid: true},
		TeamID:            sql.NullInt32{Int32: int32(CreateRandomTeam(t).ID), Valid: true},
	}

	employee2, err := testQueries.UpdateEmployee(context.Background(), args)
	require.NoError(t, err)

	require.Equal(t, employee1.ID, args.ID)
	require.Equal(t, employee2.FullName, args.FullName)

	require.NotEqual(t, employee1.FullName, employee2.FullName)

}

func TestDeleteEmployee(t *testing.T) {
	employee1 := CreateRandomEmployee(t)

	testQueries.DeleteEmployee(context.Background(), employee1.ID)

	employee2, err := testQueries.GetEmployee(context.Background(), employee1.ID)

	require.Error(t, err)
	require.Empty(t, employee2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
