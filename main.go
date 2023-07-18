package main

import (
	"crud/db"
	"crud/web/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfigs(); err != nil {
		log.Fatalf("init configs: %v", err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/user/{id}", users.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", users.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", users.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", users.DeleteUserHandler).Methods(http.MethodDelete)

	srv := http.Server{
		Addr:    viper.GetString("port"),
		Handler: router,
	}

	srv.ListenAndServe()
}

func initConfigs() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("development")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
}
