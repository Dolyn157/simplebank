package main

import (
	"database/sql"
	"log"

	api "dolyn157.dev/simplebank/api"
	"dolyn157.dev/simplebank/config"
	db "dolyn157.dev/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

func main() {
	config, err := config.LoadConfig("./config/.")
	if err != nil {
		log.Fatal("cannot load the config.", err)
	}

	conn, err := sql.Open(dbDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connected to the server.", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create the server.", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start the server.", err)
	}

}
