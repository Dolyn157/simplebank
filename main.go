package main

import (
	"database/sql"
	"log"

	api "dolyn157.dev/simplebank/api"
	db "dolyn157.dev/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:15432/simple_bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connected to the server.", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start the server.", err)
	}

}
