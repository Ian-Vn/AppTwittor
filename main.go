package main

import (
	"log"

	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/handlers"
)

func main() {
	// Checamoa la conexion a la bd, si hubo un error entonces salimos del main
	if database.ChequeoConexion() != 1 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
