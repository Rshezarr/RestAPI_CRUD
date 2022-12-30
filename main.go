package main

import (
	"crud/db"
	"crud/pkg/user"
	"log"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalln(err)
	}

	user := user.User{
		// ID:        1,
		FirstName: "test_first_name_01",
		LastName:  "test_last_name_01",
	}

	if err := user.DeleteUserByID(); err != nil {
		log.Fatalln(err)
	}
}
