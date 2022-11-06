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
	r.HandleFunc("/user/{id}", users.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", users.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", users.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", users.DeleteUserHandler).Methods(http.MethodDelete)

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	s.ListenAndServe()
}
