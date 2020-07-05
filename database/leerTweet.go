package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Funcion para leer los tweets a partir de un ID, se devuelve un slice de structs y un booleano
func LeerTweets( ID string, pagina int64 ) ( [] *models.DevovlerTweets, bool ) {
	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("tweet")

	var resultados [] *models.DevovlerTweets

	// Realizamos la condicion
	condicion := bson.M{ "userid": ID }

	// realizamos las opciones creandp un apuntador a FindOptions
	opciones := options.Find()
	// Modificamos algunas propiedades
	// para traer 20 registros
	opciones.SetLimit(20)
	// Ordenamos los documentos por el campo fecha en orden descendente
	opciones.SetSort( bson.D{ { Key: "fecha", Value: -1 }  } )
	// opcion para crear paginacion
	// Lo que se hace es contar desde eñ registro que empieza en la variable pagina y hacer la operacion
	// por ejemplo si pagina = 1 entonces la paginacion es 0 por lo cual empieza ene l registro 0 y toma 20 registros
	// si pagina = 2 entonces la paginacion es 20 y toma 20 registros a partir de la posicion 20 por lo cual deberia iniciar la siguiente
	// paginacion en el registro numero 40, y si pagina = 3 entonces la paginacion es 40
	opciones.SetSkip( (pagina - 1) * 20 )

	// Llamamos a la funcion Find para devolver mas de un registro, es decir, un apuntador a cursor
	cursor, err := col.Find(ctx, condicion, opciones)

	// si ha ocurrido un error mientras se buscaban los registros, si es vacio se devuelve un cursor vacio
	if err != nil {
		log.Fatal( err.Error() )
		return resultados, false
	}

	// si no ha ocurrido un error se recorren los documentos
	for cursor.Next( context.TODO() ) {

		// Creamos una variable del tipo devolverTweets
		var registro models.DevovlerTweets
		// se decodifica el documento hacia la estrcutura
		err := cursor.Decode( &registro )
		// si no ha ocurrido un error mientras se decodificaba hacia la estructura
		if err != nil {
			// se regresa un resultado vacio y un false
			return resultados, false
		}
		// si no ha ocurrido un error entonces se agregan elementos al slice de resultados
		resultados = append( resultados, &registro )
	}
	return resultados, true

}

