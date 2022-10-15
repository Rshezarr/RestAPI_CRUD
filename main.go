package main

import (
	"crud/db"
	"log"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalln(err)
	}
}
