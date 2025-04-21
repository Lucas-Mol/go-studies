package config

import (
	"database/sql"
	"log"
)

func InitializeDBConn(dbURL string) *sql.DB {
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}
	return conn
}
