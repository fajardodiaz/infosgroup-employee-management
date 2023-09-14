package main

import (
	"database/sql"
	"log"

	"github.com/fajardodiaz/infosgroup-employee-management/internal/api"
	"github.com/fajardodiaz/infosgroup-employee-management/utils"
	_ "github.com/lib/pq"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
