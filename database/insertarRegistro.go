package database

import (
	"context"
	"time"

	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertarRegistro es una funcion que inserta un registro en la coleccion*/
func InsertarRegistro(usuario models.Usuario) (string, bool, error) {

	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("usuarios")

	// Encriptamos el password
	usuario.Password, _ = EncriptarPassword(usuario.Password)

	// Insertamos el registro
	result, err := col.InsertOne(ctx, usuario)

	// Si ha ocurrido un error
	if err != nil {
		return "", false, err
	}

	// Si no ha ocurrido el error obtenemos el valor del ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	// Convertimmos el ObjID a un string
	return ObjID.String(), true, nil

}
