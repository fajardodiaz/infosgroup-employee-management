package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	_ "github.com/lib/pq"
)

var testQueries *database.Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:8082/infosgroup-employee-management-db?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = database.New(testDB)

	os.Exit(m.Run())
}
