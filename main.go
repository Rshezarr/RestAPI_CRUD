package main

import (
	"log"

	_ "github.com/lib/pq"

	"crud/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
