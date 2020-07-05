package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Funcion para leer todos los usuarios
func LeerTodosUusuarios( ID string, page int64, search string, tipo string ) ([] *models.Usuario, bool)  {

	// Creamos un contexto para terminar el 'contexto' en un tiempo
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// Liberar los recursos del contexto
	defer cancel()

	// Hacemos la conexion a la base de datos de Mongo
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("usuarios")

	// creamos un slice de tipo usuario
	var resultados [] *models.Usuario

	// Creamos un conjunto de opciones
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit( 20 )

	// Realizamos una condicion
	query := bson.M{ "nombre": bson.M{ "$regex": `(?i)` + search  }}

	// Creamos un conjunto de documentos con la funcion Find
	cur, err := col.Find(ctx, query, findOptions)

	// Si ha ocurrido un error es decir que no haya documentos se devuelve false y resultados vacios
	if err != nil {
		return resultados, false
	}

	// si no ha ocurrido un error entonces recorremos el cursos
	var encontrado, incluir bool

	//f recorremos el cursor
	for cur.Next(ctx) {
		// creamos una variable del tipo usuarios
		var usuario models.Usuario
		// decodificamos el documento hacia la variable usuarios
		err := cur.Decode( &usuario )
		// su ha ocurrido un error
		if err != nil {
			// se devuelve el slice vacio y un false
			return resultados, false
		}

		// se crea una variable del tipo relacion
		var relacion models.Relacion
		// se modifican algunas propiedades
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = usuario.ID.Hex()

		// setemos una variable false
		incluir = false

		// Se verifica la relacion
		encontrado, err = ConsultarRelacion( relacion )

		// se verifica el tipo
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		// si el ID es el mismo
		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}

		// la incluision es true
		if incluir {
			usuario.Password = ""
			usuario.Biografia = ""
			usuario.SitioWeb = ""
			usuario.Ubicacion = ""
			usuario.Banner = ""
			usuario.Email = ""
			// Se agrega a ña lista
			resultados = append( resultados, &usuario )
		}
	}

	// si ha ocurrido un error
	err = cur.Err()
	if err != nil {
		return resultados, false
	}
	// si no ha habido error
	// cerramos el cursor
	cur.Close(ctx)
	return resultados, true

}

