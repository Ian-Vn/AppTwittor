package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Funcion pra leer los tweets de los seguidores
func LeerTweetsSeguidores(ID string, pagina int) ([] models.DevolverTweetSeguidores, bool)  {
	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("relacion")

	// se crea la pagina
	skip := (pagina - 1 ) * 20

	// se crea la condicion
	condiciones := make( []bson.M, 0 )
	condiciones = append( condiciones, bson.M{ "$match": bson.M{ "usuarioid": ID } } )
	condiciones = append( condiciones, bson.M{ "$lookup": bson.M{ "from": "tweet", "localField": "usuariorelacionid", "foreignField": "userid", "as": "tweet" }})
	condiciones = append(condiciones, bson.M{ "$unwind": "$tweet" })
	condiciones = append(condiciones, bson.M{"$sort": bson.M{ "tweet.fecha": -1 }})
	condiciones = append(condiciones, bson.M{ "$skip": skip })
	condiciones = append(condiciones, bson.M{ "$limit": 20 })

	// Se usa la funcion agregate
	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevolverTweetSeguidores

	// Se usa decodifcan los documentos al slice
	err = cursor.All(ctx, &result)

	// si ha ocurrido un erro
	if err != nil {
		return result, false
	}
	return result, true
}

