package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Funcion para borrar un tweet, se toma el id del tweet y el del usuario
func BorrarTweet( ID string, UserID string ) error {
	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("tweet")

	// Convertimos el ID a un ObjectID
	objID, _ := primitive.ObjectIDFromHex( ID )

	// Creamos la condicion
	//condicion := bson.M{ "_id": objID, "userid": UserID  }
	condicion := bson.D{ {"_id", objID}, { "userid", UserID }   }

	// Llamamops a la funcion deleteone
	_, err := col.DeleteOne( ctx, condicion )

	// solo regreamos el error
	return err

}

