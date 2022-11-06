package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	if err := InitConfigs(); err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}

	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("configs.host"),     //host
		viper.GetString("configs.port"),     //port
		viper.GetString("configs.username"), // user
		viper.GetString("configs.dbname"),   //dbname
		viper.GetString("configs.password"), //password
		viper.GetString("configs.sslmode"),  //sslmode
	)

	DB, err = sqlx.Connect(viper.GetString("configs.driver"), conn)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	if goose.Up(DB.DB, "./migration", goose.WithNoVersioning()); err != nil {
		log.Fatalf("Err: migr up %v", err)
	}

	// switch os.Args[1] {
	// case "up":

	// case "down":
	// 	if err := goose.Down(DB.DB, "./migration"); err != nil {
	// 		log.Fatalf("Err: migr down %v", err)
	// 	}
	// default:
	// 	log.Fatalln("not enough args")
	// }
	return nil
}

func InitConfigs() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("development")
	return viper.ReadInConfig()
}
