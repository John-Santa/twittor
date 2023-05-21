package main

import (
	"log"

	"github.com/John-Santa/twittor/db"
	"github.com/John-Santa/twittor/handlers"
)

func main() {
	if !db.HaveConnection() {
		log.Fatal("Sin conexion a DB")
	}
	handlers.Routes()
}
