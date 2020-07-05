package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Funcion para insertar un tweet
func InsertarTweet( Tweet models.TweetGrabar ) (string, bool, error) {

	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("tweet")

	// Creamos un bson
	registro := bson.M{
		"userid": Tweet.UserID,
		"mensaje": Tweet.Mensaje,
		"fecha": Tweet.Fecha,
	}

	// Se graba el resultado con el metodo InsertOne
	result, err := col.InsertOne(ctx, registro )

	// si ha ocurrido un error
	if err!= nil {
		return "", false ,err
	}
	// si no hay error regresamos el objectID del documento
	objID, _ := result.InsertedID.(primitive.ObjectID)

	// regresaos el id
	return objID.String(), true, nil
}
