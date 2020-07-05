package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"time"
)

// Funcion para borrar una relacion
func BorrarRelacion( relacion models.Relacion) (bool, error) {
	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("relacion")

	// podemos pasar una estructura ya que es clave-valor y eso se interpeta como una condicion
	_, err := col.DeleteOne(ctx, relacion)

	// si ha ocurrido un error
	if err != nil {
		return false, err
	}

	// si no ha ocurrido un error
	return true, nil

}

