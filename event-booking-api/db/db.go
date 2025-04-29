package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "event-booking.db")
	if err != nil {
		panic("Could not connect to DB: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	runMigrations()
}

func runMigrations() {
	createUsersTable()
	createEventsTable()
	createRegistrationTable()
}

func createUsersTable() {
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS tb_users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    email TEXT NOT NULL UNIQUE,
	    password TEXT NOT NULL
	)
	`

	execMigrations("tb_users", createUsersTableQuery)
}

func createEventsTable() {
	createEventsTableQuery := `
	CREATE TABLE IF NOT EXISTS tb_events (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT NOT NULL,
	    description TEXT NOT NULL,
	    location TEXT NOT NULL,
	    date_time DATETIME NOT NULL,
	    user_id INTEGER,
	     FOREIGN KEY (user_id) REFERENCES tb_users(id)
	)
	`

	execMigrations("tb_events", createEventsTableQuery)
}

func createRegistrationTable() {
	createRegistrationTableQuery := `
	CREATE TABLE IF NOT EXISTS tb_registrations (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
	    user_id INTEGER NOT NULL,
	     	FOREIGN KEY (event_id) REFERENCES tb_events(id),
	    	FOREIGN KEY (user_id) REFERENCES tb_users(id)
	)
	`

	execMigrations("tb_registrations", createRegistrationTableQuery)
}

func execMigrations(tableName, query string) {
	_, err := DB.Exec(query)
	if err != nil {
		panic(fmt.Sprintf("Could not create %s: %s", tableName, err.Error()))
	}
}
