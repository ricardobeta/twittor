package main

import (
	"log"

	"github.com/ricardobeta/twittor/bd"
	"github.com/ricardobeta/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexión a la BD")
		return
	}

	handlers.Manejadores()
}
