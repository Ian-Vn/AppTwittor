package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// Funcion que busca un perfil
func BuscarPerfil(ID string) (models.Usuario, error) {
	// Creamos un contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Nuevamente usamos el cliente para conectarnos a la bd
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("usuarios")

	// Variable para grabar los datos
	var perfil models.Usuario

	// Obtenemos el tipo de dato del objectID a partir del string
	objID, _ := primitive.ObjectIDFromHex(ID)

	// Creamos la condicion
	condicion := bson.M{ "_id": objID }

	// Hacemos la consulta
	err := col.FindOne(ctx, condicion).Decode(&perfil)

	// Si ha ocurrido un error
	if err!= nil {
		log.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
