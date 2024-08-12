package main

import (
	"database/sql"
	"fmt"
	"log"
	api "simplebank/api"
	db "simplebank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:postgres@123@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	fmt.Println(serverAddress)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}
}
