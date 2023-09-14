package tests

import (
	"database/sql"
	"log"
	"os"
	"testing"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/fajardodiaz/infosgroup-employee-management/utils"
	_ "github.com/lib/pq"
)

var testQueries *database.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../../..")
	if err != nil {
		log.Fatal("cannor read config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = database.New(testDB)

	os.Exit(m.Run())
}
