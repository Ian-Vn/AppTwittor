package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Funcion para consultar las relaciones
func ConsultarRelacion( relacion models.Relacion) (bool, error) {
	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("relacion")

	// creamos una condicion
	condicion := bson.D{ { "usuarioid", relacion.UsuarioID }, { "usuariorelacionid", relacion.UsuarioRelacionID } }

	// creamos una variable para almacenar el resultado
	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode( &resultado )

	// si ha ocurrido un error es decir si no encontro un resultado entonces se mandara un error
	if err != nil {
		return false, err
	}
	return true, nil
}
