package database

import (
	"context"
	"time"

	"github.com/Ian-Vn/AppTwittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ExisteUsuario recibe un email y verifica en la bd si existe dicho usuario*/
func ExisteUsuario(email string) (models.Usuario, bool, string) {

	// Creamos un contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Nuevamente usamos el cliente para conectarnos a la bd
	db := Cliente.Database("twitter")
	// Nos conectamos a la coleccion
	col := db.Collection("usuarios")

	// Creamos un bson a partir de un mapa string,interface{} para hacer una sola condicion
	condicion := bson.M{ "email": email }

	// Creamos un modelo para almacenar el objeto encontrado de la bd
	var resultado models.Usuario

	// Buscamos el registro pasandole el contexto y la condicion y lo decodificamos al struct
	// con el findone si no hay coincidencias se devuelve un error y por ende sus metodos devolveran error
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	// Se obtiene el ID del usuario, si no hay se obtiene un string vacio
	ID := resultado.ID.Hex()

	// Si hay un error entonces se devulve la estructura vacia, un false y el ID que es un string vacio
	if err != nil {
		return resultado, false, ID
	}
	// Si no hay error se devulve la estructura, un true y el ID
	return resultado, true, ID

}
