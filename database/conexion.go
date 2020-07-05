package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Cliente es una variable con un apuntador a un client*/
var Cliente = ConectarBD()

// Creando un cliente, en el cual le pasammos la URL de la BD
var clienteOpcion = options.Client().ApplyURI("mongodb+srv://ian:HK12F_99@twittor-cu78n.mongodb.net/test?retryWrites=true&w=majority")

/*ConectarBD es la funcion que devuelve un apuntador a client que sera usado para volver a hacer ping a la bd */
func ConectarBD() *mongo.Client {

	// Creamos un cliente
	cliente, err := mongo.Connect(context.TODO(), clienteOpcion)

	// checamos si ha habido un error durante la conexion
	if err != nil {
		log.Fatal(err.Error())
		return cliente
	}

	// Como client devuelve un apuntador a *client entonces podemos usar el metodo para poder verificar que el cliente ha podido
	// conectarse a la bd
	err = cliente.Ping(context.TODO(), nil)

	// si ha ocurrido un error entonces devolvemos el cliente
	if err != nil {
		log.Fatal(err.Error())
		return cliente
	}

	// Si todo ha salido bien devolvemos el cliente
	log.Println("Conexión exitosa a la BD")
	return cliente

}

/*ChequeoConexion es la funcion que toma el valor de la variable cliente y realizar un ping de nuevo */
func ChequeoConexion() int {
	// volvemos a verificar la conexion pero ahora con la variable signada a la funcion anterior
	err := Cliente.Ping(context.TODO(), nil)
	// Si ha ocurrido un error entonces devolvemos 0 de lo contrario 1
	if err != nil {
		return 0
	}
	return 1
}
