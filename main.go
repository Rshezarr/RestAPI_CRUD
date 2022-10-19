package main

import (
	"crud/db"
	"crud/web/users"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalln(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/user/", users.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/", users.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/", users.UpdateUserHeader).Methods(http.MethodPut)

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	s.ListenAndServe()
}
