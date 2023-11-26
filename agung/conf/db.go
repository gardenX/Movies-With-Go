package conf

import (
	"database/sql"
	"os"
)

var (
	Db *sql.DB
)

func Connect() (err error) {
	Db, err = sql.Open("postgres", os.Getenv("PSQL_DSN"))

	return
}

func Close() (err error) {
	err = Db.Close()

	return
}
