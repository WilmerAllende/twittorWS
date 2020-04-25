package main

import (
	"log"
	"github.com/WilmerAllende/twittorWS/handlers"
	"github.com/WilmerAllende/twittorWS/bd"
)

func main()  {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("SIN CONEXION A LA BD")
		return
	}
	handlers.Manejadores()
}