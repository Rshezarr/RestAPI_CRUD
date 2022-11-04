package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const (
	userTable = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		data VARCHAR
	);`
)

func InitDB() error {
	if err := InitConfigs(); err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("services.postgres.configs.0"), //host
		viper.GetString("services.postgres.expose.0"),  //port
		viper.GetString("services.postgres.configs.1"), // user
		viper.GetString("services.postgres.configs.1"), //dbname
		viper.GetString("services.postgres.configs.2"), //password
		viper.GetString("services.postgres.configs.3"), //sslmode
	)

	fmt.Println(conn)

	DB, err = sqlx.Connect(viper.GetString("services.postgres.container_name"), conn)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
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

func InitConfigs() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("docker-compose")
	return viper.ReadInConfig()
}
