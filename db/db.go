package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	var err error

	DB, err = sqlx.Connect(viper.GetString("dbDriver"), viper.GetString("databaseURL"))

	if err != nil {
		return fmt.Errorf("database connection %w", err)
	}

	if goose.Up(DB.DB, "./migration", goose.WithNoVersioning()); err != nil {
		return fmt.Errorf("error: migration up %w", err)
	}

	// if goose.Down(DB.DB, "./migration", goose.WithNoVersioning()); err != nil {
	// 	return fmt.Errorf("error: migration up %w", err)
	// }

	return nil
}
