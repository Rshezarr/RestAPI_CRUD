package db

import (
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const (
	driver      = "postgres"
	databaseUrl = "host=localhost port=5432 user=postgres dbname=postgres password=qwerty sslmode=disable"

	userTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50)
		);`
)

func InitDB() error {
	var err error
	DB, err = sqlx.Open(driver, databaseUrl)
	if err != nil {
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	_, err = DB.Exec(userTable)
	if err != nil {
		return err
	}

	return nil
}
