package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var Db *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	Db, err = sql.Open("sqlite", "./api.db")
	if err != nil {
		panic(err)
	}

	Db.SetMaxOpenConns(17)
	Db.SetMaxIdleConns(7)

	createTableSQL()
	createUsersSQL()
	createRegistrationsTable()

	return Db, nil
}

func createTableSQL() {

	createEventsSQL := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	if _, err := Db.Exec(createEventsSQL); err != nil {
		panic("Couldn't create events table: " + err.Error())
	}
}

func createUsersSQL() {
	createUsersSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	if _, err := Db.Exec(createUsersSQL); err != nil {
		panic("Couldn't create users table: " + err.Error())
	}
}

func createRegistrationsTable() {
	createRegistrationsSQL := `CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	if _, err := Db.Exec(createRegistrationsSQL); err != nil {
		panic("Couldn't create registrations table: " + err.Error())
	}
}
