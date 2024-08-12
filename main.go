package main

import (
	"database/sql"
	"fmt"
	"log"
	api "simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".") // . because the config file is in same location as main.go file
	if err != nil {
		log.Fatal("Can not load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	fmt.Println(config.ServerAddress)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}
}
