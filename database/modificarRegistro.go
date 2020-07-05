package database

import (
	"context"
	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Funcion para modificar el registro, el cual recibe el id para modificar el registro a partir de este id
func ModificarRegistro(usuario models.Usuario, ID string) (bool, error) {

	// Creamos un contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Nuevamente usamos el cliente para conectarnos a la bd
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("usuarios")

	// Creamos un mapa
	registro := make( map[string]interface{})

	// Agregamos algunas claves
	if len(usuario.Nombre) > 0 {
		registro["nombre"] = usuario.Nombre
	}
	if len(usuario.Apellidos) > 0 {
		registro["apellidos"] = usuario.Apellidos
	}

	registro["fechaNacimiento"] = usuario.FechaNacimiento

	if len(usuario.Avatar) > 0 {
		registro["avatar"] = usuario.Avatar
	}

	if len(usuario.Banner) > 0 {
		registro["banner"] = usuario.Banner
	}

	if len(usuario.Biografia) > 0 {
		registro["biografia"] = usuario.Biografia
	}

	if len(usuario.Ubicacion) > 0 {
		registro["ubicacion"] = usuario.Ubicacion
	}

	if len(usuario.SitioWeb) > 0 {
		registro["sitioWeb"] = usuario.SitioWeb
	}

	// Usamos la bandera para indicar acutalizacion, esto indica que se actualizaran las claves
	// que estan en el mapa de registro
	updtString := bson.M{ "$set": registro }

	// Convertimos el id a un tipo primitive objectid usando la represebtacuib hexadecimal
	objID, _ := primitive.ObjectIDFromHex( ID )

	// creamos el filtro en donde indicamos que el id debe ser igual al ID que creamos del objID
	filtro := bson.M{ "_id": bson.M{ "$eq": objID} }

	// llamamos a la funcion de actualizar
	_, err := col.UpdateOne(ctx, filtro, updtString )

	// si ha ocurrido un error
	if err != nil {
		return false, err
	}

	// si no ha ocurrido un error
	return true, nil
}



