package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	connection := "user=pgadmin dbname=testdb password=secure_password host=127.0.0.1 sslmode=disable" // test parameters
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
