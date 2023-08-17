package main

import (
	"database/sql"
	"log"

	"github.com/fajardodiaz/infosgroup-employee-management/internal/api"
	_ "github.com/lib/pq"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:8082/infosgroup-employee-management-db?sslmode=disable"
	serverAddress = "0.0.0.0:9000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal(err)
	}

}
