package routes

import (
	"github.com/Ian-Vn/AppTwittor/database"
	"github.com/Ian-Vn/AppTwittor/models"
	"io"
	"net/http"
	"os"
	"strings"
)

// Funcion para subir un avatar
func SubirAvatar( w http.ResponseWriter, r *http.Request ) {

	// llamamos al metodo FormFile en el cual la clave va a ser avatar, el primer dato es una interface, el segundo es un struct con informacion
	// de archivo
	file, handler, err := r.FormFile("avatar")

	// obtenemos el nombre del archivo llamando a split el cual crea un slice con el nombre del archivo
	// y la extension
	var extension = strings.Split( handler.Filename, "." )[1]

	// establecemos el nombre del archivo
	var archivo = "uploads/avatars/" + IDUsuario + "." + extension

	// Creamos un nuevo archivo, este no es como tal el archivo sino como uno "temporal"
	// las banderas indican que se abre en modo de lectura/escritura y si no existe lp crea
	f, err := os.OpenFile(archivo, os.O_WRONLY | os.O_CREATE, 0666 )

	// si ha ocurrido un error mandamos un mensaje de error
	if err != nil {
		http.Error(w, "Error al subir la imagen "+ err.Error(), http.StatusBadRequest)
		return
	}

	// Reemplazamos el archivo hacia del request hacia el destino
	_, err = io.Copy(f, file)

	// si ha ocurrido un error mientras se copiaba el archivo
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+ err.Error(), http.StatusBadRequest)
		return
	}

	// Creamos una variable del tipo usuario
	var usuario models.Usuario
	// Creamos un status
	var status bool

	// modificamos la propiedad del avatar
	usuario.Avatar = IDUsuario + "." + extension
	// llamamos a la funcion de modificar registro en donde se le pasa un usuario y el id
	// primero se busca el documento referente al ID y se crea un mapa con las opciones a modiciar
	// es decir se verifican si algunas propiedades del modelo de usuario viene vacia, si no viene vacia
	// se crea la clave-valor y son los datos a actualizat
	status, err = database.ModificarRegistro( usuario, IDUsuario )
	// si ha ocurrido un error o el status es false
	if err != nil || !status {
		http.Error(w, "Error grabar el avatar en la BD "+ err.Error(), http.StatusBadRequest)
		return
	}

	// setemos el header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

