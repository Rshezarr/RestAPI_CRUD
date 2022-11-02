package db

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const (
	Host     = "localhost"
	Port     = "5436"
	Username = "postgres"
	Password = "qwerty"
	DBName   = "postgres"
	SSLmode  = "disable"

	userTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		data VARCHAR
		);`
)

func InitDB() error {
	var err error
	DB, err = sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Host, Port, Username, DBName, Password, SSLmode))
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
